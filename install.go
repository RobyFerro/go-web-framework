package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
)

// Install method will install Go-Web
type Install struct {
	Signature   string
	Description string
	Args        string
}

// Register this command
func (c *Install) Register() {
	c.Signature = "install"          // Change command signature
	c.Description = "install Go-Web" // Change command description
}

// Run this command
func (c *Install) Run() {
	var _, filename, _, _ = runtime.Caller(0)
	if err := dir(filepath.Join(path.Dir(filename), "../../"), c.Args); err != nil {
		ProcessError(err)
	}
}

// Dir copies a whole directory recursively
func dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = file(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// File copies a single file from src to dst
func file(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err := os.Chown(dst, syscall.Getuid(), syscall.Getgid()); err != nil {
		ProcessError(err)
	}

	return os.Chmod(dst, srcinfo.Mode())
}
