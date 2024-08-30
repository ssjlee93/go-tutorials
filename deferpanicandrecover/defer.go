package deferpanicandrecover

import (
	"io"
	"os"
)

// CopyFileWithoutDefer demonstrates a case without defer
func CopyFileWithoutDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// if line below fails,
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	// files will never close on failure
	dst.Close()
	src.Close()
	return
}

// CopyFile opens two files and copies contents.
// demonstrates usage of §defer§ statements
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// ensure files will always close
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
