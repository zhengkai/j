package j_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
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

	t.Run(`echo`, testEcho)
	t.Run(`color`, testColor)
	t.Run(`pointless`, testPointless)
	t.Run(`file`, testFile)
	t.Run(`lineFunc`, testLineFunc)
}
