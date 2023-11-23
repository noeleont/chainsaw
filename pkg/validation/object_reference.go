package validation

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateObjectReference(path *field.Path, obj v1alpha1.ObjectReference) field.ErrorList {
	var errs field.ErrorList
	if obj.Kind == "" {
		errs = append(errs, field.Invalid(path.Child("kind"), obj, "kind must be specified"))
	}
	if obj.APIVersion == "" {
		errs = append(errs, field.Invalid(path.Child("apiVersion"), obj, "apiVersion must be specified"))
	}
	return errs
}
