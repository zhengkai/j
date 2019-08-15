package j_test

import (
	"fmt"
	"testing"

	"github.com/zhengkai/j"
)

func TestBlock(t *testing.T) {

	b := make(chan bool)
	lineFn := func(line *string) {
		b <- true
	}

	errReport := false
	errFn := func(o *j.Logger) {
		errReport = true
	}

	x := j.NewPure(&j.Config{
		Echo:    true,
		LineFn:  lineFn,
		Tunnel:  3,
		ErrorFn: errFn,
	})

	c := newCapturer()

	x.Log(`foo`)
	x.Log(`foo`)
	x.Log(`foo`)

	if x.Error != nil || errReport {
		t.Error(`block error`, x.Error)
	}

	x.Log(`foo`)
	if x.Error != j.ErrTunnelOverflow || !errReport {
		fmt.Println(`block no overflow error`, x.Error)
	}
	s := c.end()

	if s != `` {
		t.Error(`block output`, len(s))
	}
}
