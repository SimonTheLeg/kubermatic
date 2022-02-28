package validationtest

import (
	"fmt"
	"io"
	"os"
	"testing"

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

func TestValidation(t *testing.T) {

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
