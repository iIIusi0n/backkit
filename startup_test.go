package backkit_test

import (
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestHkcuRunStartup(t *testing.T) {
	if err := backkit.AddCurrentFileToStartupUsingHkcuRun("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHkcuRunStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHkcuRunOnceStartup(t *testing.T) {
	if err := backkit.AddCurrentFileToStartupUsingHkcuRunOnce("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHkcuRunOnceStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHklmRunStartup(t *testing.T) {
	if err := backkit.AddCurrentFileToStartupUsingHklmRun("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHklmRunStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}

func TestHklmRunOnceStartup(t *testing.T) {
	if err := backkit.AddCurrentFileToStartupUsingHklmRunOnce("backkit_test"); err != nil {
		t.Error(err.Error())
	}
	if err := backkit.DeleteHklmRunOnceStartup("backkit_test"); err != nil {
		t.Error(err.Error())
	}
}
