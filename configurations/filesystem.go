package configurations

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// FileSystem is an interface that shows the behavior when accessing the filesystem
type fileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
	GetDirName() string
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// OsFS implements fileSystem using the local disk.
type OsFS struct{}

func (OsFS) Open(name string) (file, error) { return os.Open(name) }

func (OsFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

func (OsFS) GetDirName() string {
	_, dirname, _, _ := runtime.Caller(0)
	fmt.Println("Dirname: ", dirname)
	return dirname
}
