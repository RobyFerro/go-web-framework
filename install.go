package gwf

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
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
	fmt.Println("Creating new Go-Web service...")
	gopath := build.Default.GOPATH
	sourcePath := fmt.Sprintf("%s/%s", gopath, "src/github.com/RobyFerro/go-web")

	if err := dir(sourcePath, c.Args); err != nil {
		ProcessError(err)
	}

	// Remove default .git directory
	if err := os.RemoveAll(fmt.Sprintf("%s/.git", c.Args)); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete!")
}

// Dir copies a whole directory recursively
func dir(src string, dst string) error {
	var (
		err     error
		fds     []os.FileInfo
		srcInfo os.FileInfo
	)

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
	var (
		err      error
		srcFile  *os.File
		destFile *os.File
		srcInfo  os.FileInfo
	)

	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()

	if destFile, err = os.Create(dst); err != nil {
		return err
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if runtime.GOOS == "linux" {
		if err := os.Chown(dst, syscall.Getuid(), syscall.Getgid()); err != nil {
			ProcessError(err)
		}
	}

	return os.Chmod(dst, srcInfo.Mode())
}
