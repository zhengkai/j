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

		if o.caller != CallerNone {
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

	if m.t == msgColor || m.t == msgColorOnce {
		o.parseMsgColor(m.t, m.content[0].(string))
		return
	}
	o.buf.Reset()

	if !m.raw {
		o.parseMsgPrefix(m)
	}

	parseByMsgType(m, o.buf)

	if !m.raw {
		o.parseMsgBR(m)
	}
}

func (o *Logger) parseMsgColor(t msgType, color string) {

	if color == `0` {
		o.useColor = false
		return
	}

	o.useColor = true
	o.color = "\x1b[" + color + `m`
	if t == msgColorOnce {
		o.stopColor = true
	}
}

func (o *Logger) parseMsgPrefix(m *msg) {

	if o.useColor {
		o.buf.WriteString(o.color)
	}
	if o.usePrefix {
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
}

func (o *Logger) parseMsgBR(m *msg) {

	addedBR := m.t == msgPrintln

	if o.useColor {

		if addedBR {
			o.buf.UnreadByte()
		}

		o.buf.WriteString("\x1b[0m")

		if o.stopColor {
			o.useColor = false
		}
	}

	if !addedBR {
		o.buf.WriteRune('\n')
	}
}

func parseByMsgType(m *msg, buf *bytes.Buffer) {

	switch m.t {

	case msgPrintln:
		buf.WriteString(fmt.Sprintln(m.content...))

	case msgPrintf:
		buf.WriteString(m.content[0].(string))

	case msgPrint:
		buf.WriteString(fmt.Sprint(m.content...))

	case msgCompact, msgRaw:
		for _, v := range m.content {
			buf.WriteString(fmt.Sprint(v))
		}

	case msgColor, msgColorOnce:
	}
}
