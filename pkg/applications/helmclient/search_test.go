package helmclient

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	helmgetter "helm.sh/helm/v3/pkg/getter"
// 	helmrepo "helm.sh/helm/v3/pkg/repo"
// 	"k8s.io/cli-runtime/pkg/genericclioptions"
// )

// func TestSearch(t *testing.T) {
// 	url := "https://charts.bitnami.com/bitnami/"

// 	restclientGetter := &genericclioptions.ConfigFlags{}
// 	c, err := NewInstallationClient(context.Background(), restclientGetter, NewSettings("/tmp/helmclient"), "default", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rep, err := c.EnsureRepository(url, AuthSettings{})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ind, err := helmrepo.LoadIndexFile(rep.CachePath + "/" + rep.Config.Name + "-index.yaml")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rep.IndexFile = ind

// 	for chartname := range rep.IndexFile.Entries {
// 		fmt.Println(chartname)
// 	}
// }

// func TestSearch2(t *testing.T) {
// 	url := "https://charts.bitnami.com/bitnami/"

// 	providers := helmgetter.Providers{helmgetter.Provider{
// 		Schemes: []string{"http", "https"},
// 		New:     helmgetter.NewHTTPGetter,
// 	},
// 	}

// 	res, err := helmrepo.FindChartInRepoURL(url, "airflow", "14.0.6", "", "", "", providers)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(res)
// }
