/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package images

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"

	addonutil "k8c.io/kubermatic/v2/pkg/addon"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

func getImagesFromAddons(log logrus.FieldLogger, addonsPath string, cluster *kubermaticv1.Cluster) ([]string, error) {
	credentials := resources.Credentials{}

	addonData, err := addonutil.NewTemplateData(cluster, credentials, "", "", "", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create addon template data: %w", err)
	}

	infos, err := os.ReadDir(addonsPath)
	if err != nil {
		return nil, fmt.Errorf("unable to list addons: %w", err)
	}

	serializer := json.NewSerializerWithOptions(&json.SimpleMetaFactory{}, scheme.Scheme, scheme.Scheme, json.SerializerOptions{})
	var images []string
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		addonName := info.Name()
		addonImages, err := getImagesFromAddon(log, path.Join(addonsPath, addonName), serializer, addonData)
		if err != nil {
			return nil, fmt.Errorf("failed to get images for addon %s: %w", addonName, err)
		}
		images = append(images, addonImages...)
	}

	return images, nil
}

func getImagesFromAddon(log logrus.FieldLogger, addonPath string, decoder runtime.Decoder, data *addonutil.TemplateData) ([]string, error) {
	log = log.WithField("addon", path.Base(addonPath))
	log.Debug("Processing manifests…")

	allManifests, err := addonutil.ParseFromFolder(zap.NewNop().Sugar(), "", addonPath, data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse addon templates in %s: %w", addonPath, err)
	}

	var images []string
	for _, manifest := range allManifests {
		manifestImages, err := getImagesFromManifest(log, decoder, manifest.Content.Raw)
		if err != nil {
			return nil, err
		}
		images = append(images, manifestImages...)
	}
	return images, nil
}
