package zj_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zhengkai/zj"
)

func TestConfig(t *testing.T) {

	x := zj.Echo - 1

	s := x.String()
	s += zj.Append.String()
	s += zj.Prefix.String()
	s += zj.Tunnel.String()
	s += zj.LineFn.String()
	s += zj.ErrorFn.String()

	var ok bool

	ok = zj.SetDefault(zj.Append, true)
	if !ok {
		t.Error(`SetDefault fail`, zj.Append)
	}

	ok = zj.SetDefault(zj.Caller, zj.CallerLong)
	if !ok {
		t.Error(`SetDefault fail`, zj.Caller)
	}

	ok = zj.SetDefault(zj.Prefix, `[prefix]`)
	if !ok {
		t.Error(`SetDefault fail`, zj.Prefix)
	}

	ok = zj.SetDefault(zj.Tunnel, 100)
	if !ok {
		t.Error(`SetDefault fail`, zj.Tunnel)
	}

	ok = zj.SetDefault(zj.PermDir, os.FileMode(0775))
	if !ok {
		t.Error(`SetDefault fail`, zj.PermDir)
	}

	ok = zj.SetDefault(zj.PermDir, 0755)
	if !ok {
		t.Error(`SetDefault fail`, zj.PermDir)
	}

	i := 0
	ok = zj.SetDefault(zj.LineFn, func(line *string) {
		i++
		fmt.Println(i)
	})
	if !ok {
		t.Error(`SetDefault fail`, zj.LineFn)
	}

	ok = zj.SetDefault(zj.ErrorFn, func(o *zj.Logger) {
		fmt.Println(o.Error)
	})
	if !ok {
		t.Error(`SetDefault fail`, zj.ErrorFn)
	}

	log := zj.NewEcho()
	log.Enable(false)
	log.Enable(true)
	log.Close()

	zj.UnsetDefault(zj.Append)
	zj.UnsetDefault(zj.Tunnel)
	zj.UnsetDefault(zj.Prefix)
	zj.UnsetDefault(zj.Caller)
	zj.UnsetDefault(zj.LineFn)
	zj.UnsetDefault(zj.ErrorFn)
}
