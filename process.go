package backkit

import (
	"os/exec"
	"syscall"
)

// Run executable file in path.
// If u set hide as true, window of process will be hid.
func RunExecutable(path string, hide bool) error {
	cmd := exec.Command(path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: hide}
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
