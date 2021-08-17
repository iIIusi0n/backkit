# backkit
Golang backdoor template for pentest

## utils
- GetCurrentPath() -> (string, error) : Get current executable full path.

## persistence
- AddStartupUsingHkcuRun(path string) : Add current path to Run registry with name.
- AddStartupUsingHkcuRun(name, path string) : Add path to Run registry with name.
- AddStartupUsingHkcuRunOnce(path string) : Add current path to RunOnce registry with name. (Will be deleted after reboot.)
- AddStartupUsingHkcuRunOnce(name, path string) : Add path to RunOnce registry with name. (Will be deleted after reboot.)
