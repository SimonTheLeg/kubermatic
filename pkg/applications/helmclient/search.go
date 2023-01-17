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
	"path/filepath"

	"go.uber.org/zap"
	helmrepo "helm.sh/helm/v3/pkg/repo"
)

// TODO @vgramer: Do you have a preference if we should keep the clients in separate files or merge them all in one client.go file?

// SearchClient is a client that allows searching through registry indices.
// If you want to use it in a concurrency context, you must create several clients with different HelmSettings. Otherwise
// writing repository.xml or download index file may fails as it will be written by several threads.
type SearchClient struct {
	*DownloadClient

	auth AuthSettings

	logger *zap.SugaredLogger
}

func NewSearchClient(ctx context.Context, settings HelmSettings, auth AuthSettings, logger *zap.SugaredLogger) *SearchClient {
	return &SearchClient{
		DownloadClient: NewDownloadClient(ctx, settings, logger),
		auth:           auth,
		logger:         logger,
	}
}

// ListAllChartsForURL downloads the index for the given registry and returns
// a list of all chart names.
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
