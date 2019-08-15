package j

import (
	"os"
	"path/filepath"
	"time"
)

const (
	flagNew    = os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_TRUNC
	flagAppend = os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_APPEND
)

// SetFile ...
func (o *Logger) SetFile(f *os.File) {
	o.file = f
	o.fileSelf = false
	o.fileFn = nil
}

// GetFile ...
func (o *Logger) GetFile() *os.File {
	return o.file
}

func (o *Logger) changeFile(t *time.Time, fileFn FileFunc) {

	if t == nil {
		now := time.Now()
		t = &now
	}

	filename := fileFn(t)
	if filename == o.filePrev {
		return
	}

	var file *os.File
	var err error
	file, err = o.openFile(filename, true)
	if err != nil {
		o.triggerError(err)
		return
	}
	o.filePrev = filename

	o.file.Sync()
	o.file.Close()

	o.file = file
}

func (o *Logger) openFile(filename string, isAppend bool) (f *os.File, err error) {

	err = checkDir(filename, o.permDir)
	if err != nil {
		return
	}

	flag := flagNew
	if isAppend {
		flag = flagAppend
	}

	f, err = os.OpenFile(filename, flag, o.permFile)
	if err != nil {
		return
	}

	return
}

func checkDir(filename string, perm os.FileMode) (err error) {

	dir, _ := filepath.Split(filename)
	if dir == `` {
		return
	}

	err = os.MkdirAll(dir, perm)
	if err != nil {
		return
	}

	return
}
