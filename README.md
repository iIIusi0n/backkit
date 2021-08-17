# backkit
Golang backdoor template for pentest

## utils
- `GetCurrentPath()` : Get current executable full path.
- `CopyDirectory(src, dst string)` : Copy directory.
- `CopyFile(src, dst string)` : Copy file.
- `DeleteFile(path string)` : Delete file.
#
- `RunExecutable(path string, hide bool)` : Run exe in `path`. If u set `hide` as true, window of process will be hidden.

## persistence
- `AddStartupUsingHkcuRun(name string)` : Add current path to Run registry with `name`.
- `AddStartupUsingHkcuRun(name, path string)` : Add `path` to Run registry with `name`.
- `AddStartupUsingHklmRun(name string)` : Add current path to Run registry with `name`. **Admin priv required**
- `AddStartupUsingHklmRun(name, path string)` : Add `path` to Run registry with `name`. **Admin priv required**
#
- `AddStartupUsingHkcuRunOnce(name string)` : Add current path to RunOnce registry with `name`. (Will be deleted after reboot.)
- `AddStartupUsingHkcuRunOnce(name, path string)` : Add `path` to RunOnce registry with `name`. (Will be deleted after reboot.)
- `AddStartupUsingHklmRunOnce(name string)` : Add current path to RunOnce registry with `name`. (Will be deleted after reboot.) **Admin priv required**
- `AddStartupUsingHklmRunOnce(name, path string)` : Add `path` to RunOnce registry with `name`. (Will be deleted after reboot.) **Admin priv required**
#
- `CopyForPersistence(path string, change, hide bool)` : Copy to another directory to prepare for someone erasing the original exe. If u set `change` as true, original exe will run copied exe and stop. About `hide`, it will make copied exe hidden.
