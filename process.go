package backkit

import "os"

// Run executable file in path.
// If u set hide as true, window of process will be hid.
func RunExecutable(path string, hide bool) error {
	var attribute os.ProcAttr
	attribute.Sys.HideWindow = hide
	_, err := os.StartProcess(path, nil, &attribute)
	return err
}
