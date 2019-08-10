package j_test

import (
	"regexp"
	"testing"

	"github.com/zhengkai/j"
)

func testEcho(t *testing.T) {

	// Log

	c := newCapturer()

	x := j.NewEcho()

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

	x1, err := j.NewPure(&j.Config{
		Echo:   true,
		Prefix: `[ T1 ] `,
		Tunnel: 1000,
		Caller: j.CallerNone,
	})

	if err != nil {
		t.Error(`func "NewPure" fail`, err)
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

	x1, err = j.New(&j.Config{
		Echo:       true,
		Tunnel:     1,
		TimeFormat: `2006`,
		Caller:     j.CallerLong,
	})

	if err != nil {
		t.Error(`func "New" fail`, err)
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

	x1, err = j.NewPure(&j.Config{
		Echo:   true,
		Prefix: "\n",
	})

	if err != nil {
		t.Error(`new logger fail`, err)
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
}
