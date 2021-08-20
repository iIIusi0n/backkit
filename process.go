package backkit

import (
	"os"
	"os/exec"
	"syscall"
)

// Run executable file in path.
// If u set hide as true, window of process will be hid.
// It will return created process's pid if succeed.
func RunExecutable(path string, hide bool) (*os.Process, error) {
	cmd := exec.Command(path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: hide}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd.Process, nil
}
