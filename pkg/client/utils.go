package client

import (
	"errors"
	"fmt"

	"github.com/kyverno/pkg/ext/output/color"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func ObjectKey(obj metav1.Object) ctrlclient.ObjectKey {
	return ctrlclient.ObjectKey{
		Name:      obj.GetName(),
		Namespace: obj.GetNamespace(),
	}
}

func Name(key ctrlclient.ObjectKey) string {
	return ColouredName(key, nil)
}

func ColouredName(key ctrlclient.ObjectKey, color *color.Color) string {
	sprint := fmt.Sprint
	if color != nil {
		sprint = color.Sprint
	}
	name := key.Name
	if name == "" {
		name = "*"
	}
	name = sprint(name)
	if key.Namespace != "" {
		name = sprint(key.Namespace) + "/" + name
	}
	return name
}

func PatchObject(actual, expected runtime.Object) (runtime.Object, error) {
	if actual == nil || expected == nil {
		return nil, errors.New("actual and expected objects must not be nil")
	}
	actualMeta, err := meta.Accessor(actual)
	if err != nil {
		return nil, err
	}
	copy := expected.DeepCopyObject()
	expectedMeta, err := meta.Accessor(copy)
	if err != nil {
		return nil, err
	}
	expectedMeta.SetResourceVersion(actualMeta.GetResourceVersion())
	return copy, nil
}