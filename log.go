// Package j implements a simple logging package. Easy to use.
//
// It has more feature than standard logger pkg, for example:
//
// · log files rotation
//
// · background (tunnel) write
//
// · ANSI color pre line
//
// · temporarily disable
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

	if !o.enable {
		return
	}

	m := &msg{
		t:       t,
		content: content,
		raw:     t == msgRaw || t == msgColor || t == msgColorOnce,
	}

	if !m.raw {

		if o.caller > CallerNone {
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

	ffn := o.fileFn
	if ffn != nil {
		o.changeFile(m.time, ffn)
	}

	if !o.parseMsg(m) {
		return
	}
	s := o.buf.String()

	if !m.raw {
		lfn := o.lineFn
		if lfn != nil {
			lfn(&s)
		}
	}

	f := o.file
	if f != nil {
		_, err = f.WriteString(s)
		if err != nil {
			o.enable = false
			o.Error = err
			return
		}
	}

	if o.echo {
		fmt.Print(s)
	}

	return
}

func (o *Logger) parseMsg(m *msg) bool {

	if m.t == msgColor || m.t == msgColorOnce {
		o.parseMsgColor(m.t, m.content[0].(string))
		return false
	}
	o.buf.Reset()

	if !m.raw {
		o.parseMsgPrefix(m)
	}

	parseByMsgType(m, o.buf)

	if !m.raw {
		o.parseMsgBR(m)
	}

	return true
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

		switch o.caller {

		case CallerShort:
			_, file = filepath.Split(file)

		case CallerShorter:
			_, file = filepath.Split(file[0 : len(file)-3])
		}

		o.buf.WriteString(fmt.Sprintf(`%s:%d `, file, m.caller.line))
	}
}

func (o *Logger) parseMsgBR(m *msg) {

	addedBR := m.t == msgPrintln

	if o.useColor {

		if addedBR {
			o.buf.Truncate(o.buf.Len() - 1)
			addedBR = false
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
