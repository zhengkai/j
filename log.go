package j

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

var (
	enable = true
)

func (o *Logger) sendLog(t msgType, content ...interface{}) (err error) {

	m := &msg{
		t:       t,
		content: content,
		raw:     t == msgRaw || t == msgColor || t == msgColorOnce,
	}

	if !m.raw {

		if o.caller > 0 {
			_, file, line, ok := runtime.Caller(2)
			c := &caller{}
			if ok {
				c.file = file
				c.line = line
			} else {
				c.file = `???`
			}
			m.caller = c
		}

		if o.useTime {
			now := time.Now()
			m.time = &now
		}
	}

	if o.useTunnel {
		o.tunnel <- m
		return
	}

	return o.doLog(m)
}

func (o *Logger) bgLog() {

	for {
		msg := <-o.tunnel

		if msg.op > 0 {

			switch msg.op {

			case opClose:
				o.enable = false
				o.stopWait.Done()
				return

			case opEnable:
				o.enable = true

			case opDisable:
				o.enable = false
			}

			continue
		}

		o.doLog(msg)
	}
}

func (o *Logger) doLog(m *msg) (err error) {

	if o.fileFunc != nil {
		o.changeFile(m.time)
	}

	o.parseMsg(m)
	s := o.buf.String()

	if !m.raw && o.lineFunc != nil {
		o.lineFunc(&s)
	}

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

func (o *Logger) parseMsg(m *msg) {
	o.buf.Reset()
	if o.usePrefix && !m.raw {
		o.buf.WriteString(o.prefix)
	}
	if m.time != nil {
		o.buf.WriteString(m.time.Format(o.timeFormat))
	}

	if m.caller != nil {
		file := m.caller.file
		if o.caller == CallerShort {
			_, file = filepath.Split(file)
		}
		o.buf.WriteString(fmt.Sprintf(`%s:%d `, file, m.caller.line))
	}

	parseByMsgType(m, o.buf)

	if m.t == msgColorOnce {
		o.colorOnce = true
	} else if o.colorOnce && m.t != msgColor {
		o.colorOnce = false
		o.buf.WriteString("\x1b[0m")
	}
}

func parseByMsgType(m *msg, buf *bytes.Buffer) {

	switch m.t {

	case msgPrintln:
		buf.WriteString(fmt.Sprintln(m.content...))

	case msgPrintf:
		buf.WriteString(fmt.Sprintf(m.content[0].(string), m.content[1:]...))

	case msgPrint:
		buf.WriteString(fmt.Sprint(m.content...))

	case msgCompact, msgRaw:
		for _, v := range m.content {
			buf.WriteString(fmt.Sprint(v))
		}

	case msgColor, msgColorOnce:
		buf.WriteString("\x1b[")
		buf.WriteString(m.content[0].(string))
		buf.WriteRune('m')
	}

	if m.t != msgPrintln && !m.raw {
		buf.WriteRune('\n')
	}
}
