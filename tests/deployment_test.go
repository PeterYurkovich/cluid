package tests

import (
	"testing"

	"github.com/peteryurkovich/cluid/pkg/apply"
	"github.com/peteryurkovich/cluid/pkg/delete"
)

func TestApplyDeployment(t *testing.T) {
	err := apply.Apply("../templates/hack/openshift/config/deployment")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteDeployment(t *testing.T) {
	err := delete.Delete("../templates/hack/openshift/config/deployment")
	if err != nil {
		t.Fatal(err)
	}
}
