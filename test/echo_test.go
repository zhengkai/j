package zj_test

import (
	"regexp"
	"testing"

	"github.com/zhengkai/zj"
)

func testEcho(t *testing.T) {

	// Log

	c := newCapturer()

	x := zj.NewEcho()

	// x.Enable(false)
	// x.Log(`hide & seek`)
	// x.Enable(true)

	x.Log(`foo`, `bar`)

	s := c.end()
	x.Close()

	replaceTime(&s)
	replaceCaller(&s)

	str := "[TIME] [CALLER] foo bar\n"

	if s != str {
		t.Error(`method "Log" fail`, s, str)
	}

	// Logf

	c = newCapturer()

	x1 := zj.NewPure(&zj.Config{
		Echo:   true,
		Prefix: `[ T1 ] `,
		Tunnel: 1000,
		Caller: zj.CallerNone,
	})

	if x1.Error != nil {
		t.Error(`func "NewPure" fail`, x.Error)
	}

	x1.Enable(false)
	x1.Enable(true)

	x1.Logf(`foo: %dv%s`, 321, `zhengkai`)
	x1.Close()

	s = c.end()

	if s != "[ T1 ] foo: 321vzhengkai\n" {
		t.Error(`method "Logf" fail`)
	}

	// Print

	c = newCapturer()

	x1 = zj.New(&zj.Config{
		Echo:       true,
		Tunnel:     1,
		TimeFormat: `2006`,
		Caller:     zj.CallerLong,
	})

	if x1.Error != nil {
		t.Error(`func "New" fail`, x1.Error)
	}

	x1.Print(`print`, 123, 321, `foo`, `bar`)
	x1.Close()

	s = c.end()

	replaceCaller(&s)

	re := regexp.MustCompile(`^20[1-9]\d((/.+)*)/\[CALLER\] print123 321foobar` + "\n$")

	if !re.MatchString(s) {
		t.Error(`method "Print" fail`)
	}

	// Compact

	c = newCapturer()

	x1 = zj.NewPure(&zj.Config{
		Echo:   true,
		Prefix: "\n",
	})

	if x1.Error != nil {
		t.Error(`new logger fail`, x1.Error)
	}

	x1.Compact(`compact`, 123, 321, `foo`, `bar`)

	s = c.end()

	if s != "\ncompact123321foobar\n" {
		t.Error(`method "Compact" fail`)
	}

	// Raw
	// BR

	c = newCapturer()

	x1.Raw(`foo`, 123)
	x1.BR()
	x1.Raw(`bar`, 321)
	x1.BR()

	s = c.end()
	if s != "foo123\nbar321\n" {
		t.Error(`method "Raw" or "BR" fail`)
	}

	c = newCapturer()

	x1 = zj.NewPure(&zj.Config{
		Echo:   true,
		Caller: zj.CallerShorter,
	})

	x1.Log(`shorter`)

	s = c.end()

	re = regexp.MustCompile(`^echo_test:\d+ shorter` + "\n$")
	if !re.MatchString(s) {
		t.Error(`caller shortest fail`, s)
	}
}
