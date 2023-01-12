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

package helmclient

import (
	"context"
	"path/filepath"

	"go.uber.org/zap"
	helmrepo "helm.sh/helm/v3/pkg/repo"
)

// SearchClient is a special type of HelmClient used for searching through registries.
// It is a bit more lightweight, as it does not require an kubernetes RESTClientGetter.
type SearchClient struct {
	*DownloadClient
	auth AuthSettings
}

func NewSearchClient(ctx context.Context, settings HelmSettings, auth AuthSettings, logger *zap.SugaredLogger) *SearchClient {
	return &SearchClient{
		DownloadClient: NewDownloadClient(ctx, settings, logger),
		auth:           auth,
	}
}

func (h *SearchClient) ListAllChartsForURL(url string) ([]string, error) {
	rep, err := h.EnsureRepository(url, h.auth)
	if err != nil {
		return nil, err
	}

	ind, err := helmrepo.LoadIndexFile(filepath.Join(rep.CachePath, rep.Config.Name) + "-index.yaml")
	if err != nil {
		return nil, err
	}
	rep.IndexFile = ind

	i := 0
	chartnames := make([]string, len(rep.IndexFile.Entries))
	for chartname := range rep.IndexFile.Entries {
		chartnames[i] = chartname
		i++
	}

	return chartnames, nil
}
