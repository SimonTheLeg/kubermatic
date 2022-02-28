package validationtest

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	ext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kube-openapi/pkg/validation/validate"
)

// Because a CRD can contain multiple versions, you have to specify the desired one
// Usually you would use .GetObjectKind().GroupVersionKind().Version) on your object
func validatorFromCRD(crd io.Reader, desVer string) (*validate.SchemaValidator, error) {

	crdv1 := &extv1.CustomResourceDefinition{}
	dec := yaml.NewYAMLOrJSONDecoder(crd, 1024)
	err := dec.Decode(crdv1)
	if err != nil {
		return nil, err
	}
	crdr := &ext.CustomResourceDefinition{}
	if err := extv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crdv1, crdr, nil); err != nil {
		return nil, err
	}

	var sv *validate.SchemaValidator
	for _, ver := range crdr.Spec.Versions {
		if ver.Name == desVer {

			// if per-version validation is enabled, use it; otherwise use global validation
			if ver.Schema != nil {
				sv, _, err = validation.NewSchemaValidator(ver.Schema)
			} else {
				sv, _, err = validation.NewSchemaValidator(crdr.Spec.Validation)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	if sv == nil {
		return nil, fmt.Errorf("Could not find SchemaValidator for desired version %q", desVer)
	}

	return sv, nil
}

func TestFullClusterValidation(t *testing.T) {

	const crdpath = "cluster_crd.yaml"
	tt := map[string]struct {
		cr string
	}{
		"valid cluster":      {"valid_cluster.yaml"},
		"illegal cni plugin": {"illegal_cniplugin.yaml"},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {

			// read in the CR, as we need its groupVersion to select the right Validator
			f, err := os.Open(tc.cr)
			if err != nil {
				t.Fatal(err)
			}
			obj := &unstructured.Unstructured{}
			dec := yaml.NewYAMLOrJSONDecoder(f, 1024)
			if err := dec.Decode(obj); err != nil {
				t.Fatal(err)
			}

			// now we open the crd and create a validator for the groupversion of our cr
			f, err = os.Open(crdpath)
			if err != nil {
				t.Fatal(err)
			}
			v, err := validatorFromCRD(f, obj.GetObjectKind().GroupVersionKind().Version)
			if err != nil {
				t.Fatal(err)
			}

			// validate the object
			// This is going to print the invalid value for our "illegal cni plugin" case
			res := validation.ValidateCustomResource(nil, obj, v)
			for _, e := range res {
				fmt.Println(e.Error())
			}

		})
	}
}

// For partial validation, we cannot make use of unstructured.Unstructured{} as
// it has fields like 'Kind' set as mandatory. As a result how we use the
// validator slightly differs. In a real case, we possibly have to use a const
// version there
func TestClusterSpecOnly(t *testing.T) {
	const crdpath = "cluster_crd.yaml"
	const crdversion = "v1"
	const fieldname = "spec"
	const fclusterspec = "illegal_cniplugin_spec_only.yaml"

	f, err := os.Open(crdpath)
	if err != nil {
		t.Fatal(err)
	}
	v, err := validatorFromCRD(f, crdversion)
	if err != nil {
		t.Fatal(err)
	}

	// in order to make partial validation possible, we wrap our resource in the struct
	// that corresponds to the full CRD and later filter out any responses that do not
	// match the fieldname
	cspec := &kubermaticv1.ClusterSpec{}
	d, err := os.ReadFile(fclusterspec)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal(d, cspec)
	if err != nil {
		t.Fatal(err)
	}
	cwrap := &kubermaticv1.Cluster{}
	cwrap.Spec = *cspec

	res := validation.ValidateCustomResource(nil, cwrap, v)
	for _, e := range res {
		if strings.HasPrefix(e.Field, fieldname) {
			fmt.Println(e.Error())
		}
	}
}
