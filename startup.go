package backkit

import (
	"errors"

	"golang.org/x/sys/windows/registry"
)

// Delete key in HKCU Run registry for startup.
func DeleteHkcuRunStartup(name string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.DeleteValue(name); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

// Add executable file to startup using Run registry in HKCU.
// AddStartupUsingHkcuRun(name, path) will add file in path to registry with key name.
// You can ignore path like AddStartupUsingHkcuRun(name), then it will add current file to registry.
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
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHkcuRun(name, currentPath)
}

// Delete key in HKCU RunOnce registry for startup.
func DeleteHkcuRunOnceStartup(name string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.DeleteValue(name); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

// Add executable file to startup using RunOnce registry in HKCU.
// AddStartupUsingHkcuRunOnce(name, path) will add file in path to registry with key name.
// You can ignore path like AddStartupUsingHkcuRunOnce(name), then it will add current file to registry.
// Value in RunOnce registry will be deleted after reboot.
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
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHkcuRunOnce(name, currentPath)
}

// Delete key in HKLM Run registry for startup.
func DeleteHklmRunStartup(name string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.DeleteValue(name); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

// Add executable file to startup using Run registry in HKLM.
// AddStartupUsingHkcuRun(name, path) will add file in path to registry with key name.
// You can ignore path like AddStartupUsingHkcuRun(name), then it will add current file to registry.
// This function required administrator privileges.
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
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHklmRun(name, currentPath)
}

// Delete key in HKLM RunOnce registry for startup.
func DeleteHklmRunOnceStartup(name string) error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	if err := key.DeleteValue(name); err != nil {
		return err
	}
	if err := key.Close(); err != nil {
		return err
	}
	return nil
}

// Add executable file to startup using RunOnce registry in HKLM.
// AddStartupUsingHkcuRunOnce(name, path) will add file in path to registry with key name.
// You can ignore path like AddStartupUsingHkcuRunOnce(name), then it will add current file to registry.
// Value in RunOnce registry will be deleted after reboot.
// This function required administrator privileges.
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
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return addExternalFileToStartupUsingHklmRunOnce(name, currentPath)
}
