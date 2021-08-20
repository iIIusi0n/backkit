package backkit

import (
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

// Add external file to HKCU Run registry for startup.
func AddExternalFileToStartupUsingHkcuRun(name, path string) error {
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

// Add current file to HKCU Run registry for startup.
func AddCurrentFileToStartupUsingHkcuRun(name string) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return AddExternalFileToStartupUsingHkcuRun(name, currentPath)
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

// Add external file to HKCU RunOnce registry for startup.
// It will be deleted after reboot.
func AddExternalFileToStartupUsingHkcuRunOnce(name, path string) error {
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

// Add current file to HKCU RunOnce registry for startup.
// It will be deleted after reboot.
func AddCurrentFileToStartupUsingHkcuRunOnce(name string) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return AddExternalFileToStartupUsingHkcuRunOnce(name, currentPath)
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

// Add external file to HKLM Run registry for startup.
// It required admin privileges.
func AddExternalFileToStartupUsingHklmRun(name, path string) error {
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

// Add current file to HKLM Run registry for startup.
// It required admin privileges.
func AddCurrentFileToStartupUsingHklmRun(name string) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return AddExternalFileToStartupUsingHklmRun(name, currentPath)
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

// Add external file to HKLM RunOnce registry for startup.
// It required admin privileges.
// It will be deleted after reboot.
func AddExternalFileToStartupUsingHklmRunOnce(name, path string) error {
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

// Add current file to HKLM RunOnce registry for startup.
// It required admin privileges.
// It will be deleted after reboot.
func AddCurrentFileToStartupUsingHklmRunOnce(name string) error {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}
	return AddExternalFileToStartupUsingHklmRunOnce(name, currentPath)
}
