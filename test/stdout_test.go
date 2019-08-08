package j_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/zhengkai/j"
)

var (
	r, w, _ = os.Pipe()
)

type capturer struct {
	r      *os.File
	w      *os.File
	stdout *os.File
}

func newCapturer() (c *capturer) {
	c = &capturer{}
	c.start()
	return
}

func (c *capturer) start() {

	c.r, c.w, _ = os.Pipe()

	c.stdout = os.Stdout
	os.Stdout = c.w
}

func (c *capturer) end() string {

	c.w.Close()
	out, _ := ioutil.ReadAll(c.r)
	c.r.Close()

	os.Stdout = c.stdout

	return string(out)
}

func TestCapturer(t *testing.T) {
	c := newCapturer()
	fmt.Print(`foo`, `bar`)
	s := c.end()

	if s != `foobar` {
		t.Fatal(`capturer fail`, s)
	}
}

func TestEcho(t *testing.T) {

	// Log

	c := newCapturer()

	x := j.NewEcho()
	x.Log(`foo`, `bar`)

	s := c.end()
	x.Close()

	re := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}\.\d{3} stdout_test.go:\d+ foo bar` + "\n$")

	if !re.MatchString(s) {
		t.Error(`method "Log" fail`)
	}

	// Logf

	c = newCapturer()

	x1, err := j.New(&j.Config{
		Echo:   true,
		Prefix: `[ T1 ] `,
		Tunnel: 1000,
		Caller: j.CallerNone,
	})

	if err != nil {
		t.Error(`func "New" fail`)
	}

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
		t.Error(`func "New" fail`)
	}

	x1.Print(`print`, 123, 321, `foo`, `bar`)
	x1.Close()

	s = c.end()

	re = regexp.MustCompile(`^20[1-9]\d((/.+)*)/stdout_test.go:\d+ print123 321foobar` + "\n$")

	if !re.MatchString(s) {
		t.Error(`method "Print" fail`)
	}

	// Compact

	c = newCapturer()

	x1, err = j.New(&j.Config{
		Echo:   true,
		Prefix: "\n",
	})

	if err != nil {
		t.Error(`func "New" fail`)
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
