package backkit_test

import (
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestHkcuRunStartup(t *testing.T) {
	if err := backkit.AddStartupUsingHkcuRun("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHkcuRunStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHkcuRunOnceStartup(t *testing.T) {
	if err := backkit.AddStartupUsingHkcuRunOnce("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHkcuRunOnceStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHklmRunStartup(t *testing.T) {
	if err := backkit.AddStartupUsingHklmRun("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHklmRunStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHklmRunOnceStartup(t *testing.T) {
	if err := backkit.AddStartupUsingHklmRunOnce("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHklmRunOnceStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}
