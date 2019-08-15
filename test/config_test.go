package j_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhengkai/j"
)

func TestConfig(t *testing.T) {

	x := j.Echo - 1

	s := x.String()
	s += j.Append.String()
	s += j.Prefix.String()
	s += j.Tunnel.String()
	s += j.LineFn.String()
	s += j.ErrorFn.String()

	var ok bool

	ok = j.SetDefault(j.Append, true)
	if !ok {
		t.Error(`SetDefault fail`, j.Append)
	}

	ok = j.SetDefault(j.Caller, j.CallerLong)
	if !ok {
		t.Error(`SetDefault fail`, j.Caller)
	}

	ok = j.SetDefault(j.Prefix, `[prefix]`)
	if !ok {
		t.Error(`SetDefault fail`, j.Prefix)
	}

	ok = j.SetDefault(j.Tunnel, 100)
	if !ok {
		t.Error(`SetDefault fail`, j.Tunnel)
	}

	ok = j.SetDefault(j.PermDir, os.FileMode(0775))
	if !ok {
		t.Error(`SetDefault fail`, j.PermDir)
	}

	ok = j.SetDefault(j.PermDir, 0755)
	if !ok {
		t.Error(`SetDefault fail`, j.PermDir)
	}

	i := 0
	ok = j.SetDefault(j.LineFn, func(line *string) {
		i++
		fmt.Println(i)
	})
	if !ok {
		t.Error(`SetDefault fail`, j.LineFn)
	}

	ok = j.SetDefault(j.ErrorFn, func(o *j.Logger) {
		fmt.Println(o.Error)
	})
	if !ok {
		t.Error(`SetDefault fail`, j.ErrorFn)
	}

	log := j.NewEcho()
	log.Enable(false)
	log.Enable(true)
	log.Close()

	j.UnsetDefault(j.Append)
	j.UnsetDefault(j.Tunnel)
	j.UnsetDefault(j.Prefix)
	j.UnsetDefault(j.Caller)
	j.UnsetDefault(j.LineFn)
	j.UnsetDefault(j.ErrorFn)
}
