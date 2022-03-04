package openapi

import (
	"fmt"
	"os"
	"testing"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func TestTest(t *testing.T) {
	obj := &kubermaticv1.Cluster{}
	/* f, err := os.ReadFile("../../../validationtest/valid_cluster.yaml") */
	f, err := os.ReadFile("../../../validationtest/illegal_cniplugin.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal(f, obj)
	if err != nil {
		t.Fatal(err)
	}
	v, err := ValidatorForType(&obj.TypeMeta)
	if err != nil {
		t.Fatal(err)
	}

	res := validation.ValidateCustomResource(nil, obj, v)

	for _, e := range res {
		fmt.Println(e.Error())
	}
}

func TestValidatorForType(t *testing.T) {
	tt := map[string]struct {
		in           *metav1.TypeMeta
		expValidator bool
		expErr       bool
	}{
		"k8c.io crd": {
			&metav1.TypeMeta{Kind: "Cluster", APIVersion: "kubermatic.k8c.io/v1"},
			true,
			false,
		},
		"apps.k8c.io crd": {
			&metav1.TypeMeta{Kind: "ApplicationDefinition", APIVersion: "apps.kubermatic.k8c.io/v1"},
			true,
			false,
		},
		"invalid kind": {
			&metav1.TypeMeta{Kind: "Invalid", APIVersion: "kubermatic.k8c.io/v1"},
			false,
			true,
		},
		"invalid apiversion": {
			&metav1.TypeMeta{Kind: "Cluster", APIVersion: "Invalid"},
			false,
			true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res, err := ValidatorForType(tc.in)

			if res != nil {
				if tc.expValidator && len(res.Schema.SchemaProps.Properties) == 0 {
					t.Errorf("Root SchemaProps.Properties are empty, when they should not be")
				}
			}

			if !tc.expErr && err != nil {
				t.Errorf("Exp error to be nil, but got %q", err)
			}

		})
	}
}
