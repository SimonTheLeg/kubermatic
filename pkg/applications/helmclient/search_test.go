/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

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

package helmclient

import (
	"context"
	"path"
	"sort"
	"testing"

	"k8c.io/kubermatic/v2/pkg/applications/test"
	kubermaticlog "k8c.io/kubermatic/v2/pkg/log"
	"k8s.io/utils/strings/slices"
)

func TestListAllChartsForRegistry(t *testing.T) {
	log := kubermaticlog.New(true, kubermaticlog.FormatJSON).Sugar()
	chartArchiveDir := t.TempDir()

	chartGlobPath := path.Join(chartArchiveDir, "*.tgz")
	test.PackageChart(t, "testdata/examplechart", chartArchiveDir)
	test.PackageChart(t, "testdata/examplechart2", chartArchiveDir)

	httpRegistryUrl := test.StartHttpRegistryWithCleanup(t, chartGlobPath)

	ociRegistryUrl := test.StartOciRegistry(t, chartGlobPath)

	testCases := []struct {
		name           string
		registryUrl    string
		expectedCharts []string
	}{
		{
			name:           "Listing from HTTP registry should be successful",
			registryUrl:    httpRegistryUrl,
			expectedCharts: []string{"examplechart", "examplechart2"},
		},
		{
			name:           "Listing from OCI registry should be successful",
			registryUrl:    ociRegistryUrl,
			expectedCharts: []string{"examplechart", "examplechart2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			downloadDest := t.TempDir()
			settings := NewSettings(downloadDest)

			c := NewSearchClient(context.Background(), settings, AuthSettings{}, log)

			res, err := c.ListAllChartsForURL(tc.registryUrl)
			if err != nil {
				t.Fatal(err)
			}

			sort.Strings(res)
			sort.Strings(tc.expectedCharts)

			if !slices.Equal(res, tc.expectedCharts) {
				t.Errorf("Expected and given charts differ. got '%v', expect '%v'", res, tc.expectedCharts)
			}
		})
	}
}
