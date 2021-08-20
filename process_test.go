package backkit_test

import (
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestRunProcess(t *testing.T) {
	err := backkit.RunExecutable("C:\\Windows\\System32\\notepad.exe", false)
	if err != nil {
		t.Error(err.Error())
	}
}
