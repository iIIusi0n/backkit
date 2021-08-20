package backkit_test

import (
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestGetPublicIP(t *testing.T) {
	if publicIP, err := backkit.GetPublicIP(); err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Public IP:", publicIP)
	}
}

func TestGetPrivateIP(t *testing.T) {
	if publicIP, err := backkit.GetPrivateIP(); err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Private IP:", publicIP)
	}
}
