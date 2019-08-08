package j

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	flagNew    = os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_TRUNC
	flagAppend = os.O_CREATE | os.O_RDWR | os.O_SYNC | os.O_APPEND
)

func (o *Logger) changeFile(t *time.Time) {

	if t == nil {
		now := time.Now()
		t = &now
	}

	filename := o.fileFunc(t)
	if filename == o.filePrev {
		return
	}

	file, err := openFile(filename, true)
	if err != nil {
		o.err = err
		o.filePrev = filename
		return
	}

	o.file.Sync()
	o.file.Close()

	o.file = file
}

func openFile(filename string, isAppend bool) (f *os.File, err error) {

	err = checkDir(filename)
	if err != nil {
		return
	}

	flag := flagNew
	if isAppend {
		flag = flagAppend
	}

	f, err = os.OpenFile(filename, flag, 0644)
	if err != nil {
		return
	}

	return
}

func checkDir(filename string) (err error) {

	dir, _ := filepath.Split(filename)
	if dir == `` {
		return fmt.Errorf(`no dir %s`, filename)
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	f, err := os.Lstat(dir)
	if err != nil {
		return
	}

	mode := f.Mode()
	if !mode.IsDir() {
		return fmt.Errorf(`not a dir "%s"`, dir)
	}

	return
}
