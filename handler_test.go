package backkit_test

import (
	"fmt"
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestHandlerCommand(t *testing.T) {
	testHandler := backkit.NewHandler("localhost", 8080, backkit.CONNECTION_NORMAL)
	err := testHandler.AddCommand("print", func(args ...string) {
		fmt.Println(args[0])
	})
	if err != nil {
		t.Error(err.Error())
	}
	testHandler.Commands["print"]("print command called")
	err = testHandler.RemoveCommand("print")
	if err != nil {
		t.Error(err.Error())
	}
}
