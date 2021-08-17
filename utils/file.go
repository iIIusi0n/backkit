package utils

import (
	"io"
	"io/ioutil"
	"os"
)

func GetCurrentPath() (string, error) {
	return os.Executable()
}

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

func DeleteFile(path string) error {
	return os.Remove(path)
}
