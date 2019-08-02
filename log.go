package j

import (
	"fmt"
	"time"
)

var (
	enable = true
)

// Log ...
func (o *Logger) Log(content ...string) (err error) {

	var t *time.Time
	if o.time != TimeNone {
		ti := time.Now()
		t = &ti
	}

	if !o.useTunnel {
		return o.log(t, content...)
	}

	o.tunnel <- &msg{
		time:    t,
		content: content,
	}

	return
}

func (o *Logger) bgLog() {
	for {
		msg := <-o.tunnel
		o.log(msg.time, msg.content...)
	}
}

func (o *Logger) log(t *time.Time, content ...string) (err error) {

	if !o.enable || !enable {
		return
	}

	o.buf.Reset()
	if t != nil {
		o.buf.WriteString(o.genTime(t))
	}

	first := true
	for _, v := range content {
		if first {
			first = false
		} else {
			o.buf.WriteRune(' ')
		}
		o.buf.WriteString(v)
	}
	o.buf.WriteRune('\n')

	s := o.buf.String()

	if o.file != nil {
		_, err = o.file.WriteString(s)
		if err != nil {
			o.enable = false
			o.err = err
			return
		}
	}

	if o.echo {
		fmt.Print(s)
	}

	return
}
