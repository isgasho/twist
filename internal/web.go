// Code generated by fileb0x at "2019-05-06 10:41:59.020858 -0700 PDT m=+0.001999373" from config file "b0x.yml" DO NOT EDIT.
// modification hash(1ae57ab045df4442fbad1ca58ed2eee7.d4c528443bf26f9fe3498e9d2ebc2775)

package internal

import (
	"bytes"

	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FilePkgHTML is "pkg.html"
var FilePkgHTML = []byte("\x3c\x21\x44\x4f\x43\x54\x59\x50\x45\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x3e\x0a\x20\x20\x20\x20\x3c\x68\x65\x61\x64\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x67\x6f\x2d\x69\x6d\x70\x6f\x72\x74\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x7b\x7b\x20\x2e\x43\x61\x6e\x6f\x6e\x69\x63\x61\x6c\x20\x7d\x7d\x20\x67\x69\x74\x20\x68\x74\x74\x70\x73\x3a\x2f\x2f\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x67\x6f\x2d\x73\x6f\x75\x72\x63\x65\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x7b\x7b\x20\x2e\x43\x61\x6e\x6f\x6e\x69\x63\x61\x6c\x20\x7d\x7d\x20\x68\x74\x74\x70\x73\x3a\x2f\x2f\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x20\x68\x74\x74\x70\x73\x3a\x2f\x2f\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x2f\x74\x72\x65\x65\x2f\x6d\x61\x73\x74\x65\x72\x7b\x2f\x64\x69\x72\x7d\x20\x68\x74\x74\x70\x73\x3a\x2f\x2f\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x2f\x74\x72\x65\x65\x2f\x6d\x61\x73\x74\x65\x72\x7b\x2f\x64\x69\x72\x7d\x2f\x7b\x66\x69\x6c\x65\x7d\x23\x4c\x7b\x6c\x69\x6e\x65\x7d\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x68\x74\x74\x70\x2d\x65\x71\x75\x69\x76\x3d\x22\x72\x65\x66\x72\x65\x73\x68\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x30\x3b\x20\x75\x72\x6c\x3d\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x22\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x20\x20\x20\x20\x3c\x62\x6f\x64\x79\x3e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x4e\x6f\x74\x68\x69\x6e\x67\x20\x74\x6f\x20\x73\x65\x65\x20\x68\x65\x72\x65\x2e\x20\x50\x6c\x65\x61\x73\x65\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x7b\x7b\x20\x2e\x53\x6f\x75\x72\x63\x65\x20\x7d\x7d\x22\x3e\x6d\x6f\x76\x65\x20\x61\x6c\x6f\x6e\x67\x3c\x2f\x61\x3e\x2e\x0a\x20\x20\x20\x20\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e\x0a")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	f, err = FS.OpenFile(CTX, "pkg.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FilePkgHTML)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}