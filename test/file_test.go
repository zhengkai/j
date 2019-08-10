package j_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/zhengkai/j"
)

func testFile(t *testing.T) {

	testFilePerm(t)

	count := 0

	j.SetDefault(j.Echo, false)

	x1, err := j.New(&j.Config{
		Filename: `log-dir/new-one.txt`,
		Echo:     false,
		PermDir:  0700,
		Tunnel:   10,
	})

	if err != nil {
		t.Error(`new file logger fail`, err)
	}

	x1.Log(`new`)
	x1.Close()

	if getPerm(`log-dir`) != 0700 {
		t.Error(`config PermDir invalid`)
	}

	count = 0
	rt := time.Now()
	x2, err := j.NewFunc(func(t *time.Time) (filename string) {
		count++
		if count > 2 {
			nt := rt.Add(time.Second)
			t = &nt
		} else {
			t = &rt
		}
		filename = t.Format(`log-dir/time-20060102-150405.txt`)
		return
	})

	if err != nil {
		t.Error(`new file logger fail`, err)
	}

	x2.Log(`tick`)
	x2.Raw(`raw`)
	x2.Log(`tick`)

	count = 0

	reFile := regexp.MustCompile(`^log-dir\/time-\d{8}-\d{6}\.txt$`)
	filepath.Walk(`log-dir`, func(path string, info os.FileInfo, err error) error {
		if reFile.MatchString(path) {
			count++
		}
		return nil
	})

	if count != 2 {
		t.Error(`filename by time fail`)
	}

	x2.SetFile(x1.GetFile())

	x2.Close()

	x2, err = j.NewFunc(func(t *time.Time) (filename string) {
		return `log-dir/time-no-change.txt`
	})
	x2.Log(`tick`)
	x2.Log(`tick`)
	x2.Close()

	x3, err := j.New(&j.Config{
		Filename: `log-dir`,
	})

	if x3 != nil || err == nil {
		t.Error(`no error when create file fail`)
	}

	x3, err = j.New(&j.Config{
		Filename: `log-dir-deny/dir/new-one.txt`,
	})

	if x3 != nil || err == nil {
		t.Error(`no error when create file fail`)
	}

	count = 0
	x3, err = j.NewFunc(func(t *time.Time) (filename string) {
		count++
		if count <= 2 {
			return `log-dir/func-success.txt`
		}
		return `log-dir-deny/fail.txt`
	})
	x3.Enable(false)
	x3.Log(`tick`)
	if x3.Error != nil {
		t.Error(`unknown erro when "NewFunc()"`)
	}
	x3.Enable(true)
	x3.Log(`tick`)
	if x3.Error != nil {
		t.Error(`Enable() 1 fail`)
	}
	x3.Log(`tick`)
	if x3.Error == nil {
		t.Error(`no error when create file fail`)
	}

	x3.Close()
	x3.Close()
	x3.Enable(false)

	x4, err := j.New(&j.Config{
		Filename: `log-dir/new-fail.txt`,
		Echo:     false,
	})

	f := x4.GetFile()
	f.Close()

	x4.Print(`fail`)

	if x4.Error == nil {
		t.Error(`no error when write file fail`)
	}

	j.SetDefault(j.Echo, true)

	testFileCount(t)
}

func testFilePerm(t *testing.T) {
	j.New(&j.Config{
		Filename: `log-file`,
		PermFile: 0600,
	})
	if getPerm(`log-file`) != 0600 {
		t.Error(`config PermFile invalid`)
	}
}

func testFileCount(t *testing.T) {
	count := 0
	filepath.Walk(`log-dir`, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			count++
		}
		return nil
	})

	if count != 6 {
		t.Error(`files in log-dir are not match`, count)
	}
}
