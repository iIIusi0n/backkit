package backkit

import (
	"io"
	"io/ioutil"
	"os"
)

// Get current file path.
func GetCurrentPath() (string, error) {
	return os.Executable()
}

// Copy directory from src to dst.
func CopyDirectory(src, dst string) error {
	err := os.MkdirAll(dst, 0777)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, f := range files {
		s := src + "/" + f.Name()
		d := dst + "/" + f.Name()
		if f.IsDir() {
			err = CopyDirectory(s, d)
		} else {
			err = CopyFile(s, d)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Copy file from src to dst.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	if err != nil {
		return err
	}
	return nil
}

// Delete file from disk.
func DeleteFile(path string) error {
	return os.Remove(path)
}
