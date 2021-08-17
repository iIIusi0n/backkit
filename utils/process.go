package utils

import "os"

func RunExecutable(path string, hide bool) error {
	var attribute os.ProcAttr
	attribute.Sys.HideWindow = hide
	_, err := os.StartProcess(path, nil, &attribute)
	return err
}
