package j_test

import (
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

	var ok bool

	j.SetDefault(j.Append, true)

	j.SetDefault(j.Caller, j.CallerLong)

	j.SetDefault(j.Prefix, `[prefix]`)

	j.SetDefault(j.Tunnel, 100)

	ok = j.SetDefault(j.PermDir, os.FileMode(0775))
	if !ok {
		t.Error(`SetDefault fail`)
	}

	j.SetDefault(j.PermDir, 0755)
	if !ok {
		t.Error(`SetDefault fail`)
	}

	ok = j.SetDefault(j.PermDir, `0775`)
	if ok {
		t.Error(`SetDefault fail`)
	}

	i := 0
	j.SetDefault(j.LineFn, func(line *string) {
		i++
	})

	log := j.NewEcho()
	log.Enable(false)
	log.Enable(true)
	log.Close()

	j.UnsetDefault(j.Append)
	j.UnsetDefault(j.Tunnel)
	j.UnsetDefault(j.Prefix)
	j.UnsetDefault(j.Caller)
	j.UnsetDefault(j.LineFn)
}
