package backkit_test

import (
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestGenerateUniqueID(t *testing.T) {
	id, err := backkit.GenerateUniqueID()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(id)
}
