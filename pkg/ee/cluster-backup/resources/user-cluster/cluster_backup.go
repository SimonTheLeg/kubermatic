//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2023 Kubermatic GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package userclusterresources

import (
	"fmt"

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/kubernetes"
	"k8c.io/kubermatic/v2/pkg/resources"
	kkpreconciling "k8c.io/kubermatic/v2/pkg/resources/reconciling"
	"k8c.io/kubermatic/v2/pkg/resources/registry"
	"k8c.io/reconciler/pkg/reconciling"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	ClusterRoleBindingName     = "velero"
	clusterBackupAppName       = "velero"
	DefaultBSLName             = "default-cluster-backup-bsl"
	customizationConfigMapName = "fs-restore-action-config"

	CloudCredentialsSecretName           = "velero-cloud-credentials"
	defaultCloudCredentialsSecretKeyName = "cloud"

	version       = "v1.12.0"
	pluginVersion = "v1.2.1"
)

// NamespaceReconciler creates the namespace for velero related resources on the user cluster.
func NamespaceReconciler() reconciling.NamedNamespaceReconcilerFactory {
	return func() (string, reconciling.NamespaceReconciler) {
		return resources.ClusterBackupNamespaceName, func(ns *corev1.Namespace) (*corev1.Namespace, error) {
			// Generated by the velero installer (velero install) to enforce compatibility with Pod Security Standards and added here for upstream compatibility.
			kubernetes.EnsureLabels(ns, map[string]string{
				"pod-security.kubernetes.io/audit":           "privileged",
				"pod-security.kubernetes.io/audit-version":   "latest",
				"pod-security.kubernetes.io/enforce":         "privileged",
				"pod-security.kubernetes.io/enforce-version": "latest",
				"pod-security.kubernetes.io/warn":            "privileged",
				"pod-security.kubernetes.io/warn-version":    "latest",
			})

			return ns, nil
		}
	}
}

// ServiceAccountReconciler creates the service account for velero on the user cluster.
func ServiceAccountReconciler() reconciling.NamedServiceAccountReconcilerFactory {
	return func() (string, reconciling.ServiceAccountReconciler) {
		return resources.ClusterBackupServiceAccountName, func(sa *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
			return sa, nil
		}
	}
}

// ClusterRoleBindingReconciler creates the clusterrolebinding for velero on the user cluster.
func ClusterRoleBindingReconciler() reconciling.NamedClusterRoleBindingReconcilerFactory {
	return func() (string, reconciling.ClusterRoleBindingReconciler) {
		return ClusterRoleBindingName, func(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			kubernetes.EnsureLabels(crb, resources.BaseAppLabels(clusterBackupAppName, nil))

			crb.RoleRef = rbacv1.RoleRef{
				// too wide but probably needed to be able to do backups and restore.
				Name:     "cluster-admin",
				Kind:     "ClusterRole",
				APIGroup: rbacv1.GroupName,
			}
			crb.Subjects = []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      resources.ClusterBackupServiceAccountName,
					Namespace: resources.ClusterBackupNamespaceName,
				},
			}
			return crb, nil
		}
	}
}

// BSLReconciler creates the default BackupStorage location is created for velero.
func BSLReconciler(cluster *kubermaticv1.Cluster, cbsl *kubermaticv1.ClusterBackupStorageLocation) kkpreconciling.NamedBackupStorageLocationReconcilerFactory {
	return func() (string, kkpreconciling.BackupStorageLocationReconciler) {
		return DefaultBSLName, func(bsl *velerov1.BackupStorageLocation) (*velerov1.BackupStorageLocation, error) {
			kubernetes.EnsureLabels(bsl, resources.BaseAppLabels(clusterBackupAppName, nil))

			projectID, ok := cluster.Labels[kubermaticv1.ProjectIDLabelKey]
			if !ok {
				return nil, fmt.Errorf("cluster ProjectID label is not set")
			}
			bsl.Spec = *cbsl.Spec.DeepCopy()
			// we set this bsl as default and remove the secret reference to make it use the default velero secret.
			bsl.Spec.Default = true
			bsl.Spec.Credential = nil
			// add bucket prefix using projectID/clusterID to avoid collision.
			bsl.Spec.ObjectStorage.Prefix = fmt.Sprintf("%s/%s", projectID, cluster.Name)
			return bsl, nil
		}
	}
}

// CustomizationConfigMapReconciler reconciles a ConfigMap to configure the velero restore helper,
// see https://velero.io/docs/v1.12/file-system-backup/#customize-restore-helper-container.
func CustomizationConfigMapReconciler(rewriter registry.ImageRewriter) reconciling.NamedConfigMapReconcilerFactory {
	return func() (string, reconciling.ConfigMapReconciler) {
		return customizationConfigMapName, func(cm *corev1.ConfigMap) (*corev1.ConfigMap, error) {
			// ensure Velero can find this ConfigMap
			kubernetes.EnsureLabels(cm, map[string]string{
				"velero.io/plugin-config":      "",
				"velero.io/pod-volume-restore": "RestoreItemAction",
			})

			if cm.Data == nil {
				cm.Data = map[string]string{}
			}

			// all we want to set is the custom restore helper image
			rewritten, err := rewriter("docker.io/velero/velero-restore-helper")
			if err != nil {
				return nil, fmt.Errorf("failed to rewrite image: %w", err)
			}

			cm.Data["image"] = rewritten

			return cm, nil
		}
	}
}
