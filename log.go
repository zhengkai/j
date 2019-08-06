package j

import (
	"fmt"
	"time"
)

var (
	enable = true
)

// Log just like log.Println
func (o *Logger) Log(a ...interface{}) (err error) {
	return o.sendLog(msgPrintln, a...)
}

// Logf just like log.Printf
func (o *Logger) Logf(format string, a ...interface{}) (err error) {
	return o.sendLog(msgPrintf, format, fmt.Sprintf(format, a...))
}

// Print just like log.Print, added linebreak
func (o *Logger) Print(format string, a ...interface{}) (err error) {
	return o.sendLog(msgPrint, a...)
}

// Compact just like fmt.Print, but no spaces
func (o *Logger) Compact(format string, a ...interface{}) (err error) {
	return o.sendLog(msgCompact, format, fmt.Sprintf(format, a...))
}

// Raw log raw (no time, no linebreak, no spaces)
func (o *Logger) Raw(a ...interface{}) (err error) {
	return o.sendLog(msgRaw, a...)
}

// BR add an empty line
func (o *Logger) BR() (err error) {
	return o.sendLog(msgRaw, "\n")
}

func (o *Logger) sendLog(t msgType, content ...interface{}) (err error) {

	m := &msg{
		t:       t,
		content: content,
		raw:     t == msgRaw || t == msgColor,
	}
	if o.useTime && !m.raw {
		now := time.Now()
		m.time = &now
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

	o.buf.Reset()
	if o.usePrefix && !m.raw {
		o.buf.WriteString(o.prefix)
	}
	if m.time != nil {
		o.buf.WriteString(m.time.Format(o.timeFormat))
	}

	switch m.t {
	case msgPrintln:
		o.buf.WriteString(fmt.Sprintln(m.content...))
	case msgPrintf:
		o.buf.WriteString(fmt.Sprintf(m.content[0].(string), m.content[1:]...))
	case msgPrint:
		o.buf.WriteString(fmt.Sprint(m.content...))
	case msgCompact, msgRaw:
		for _, v := range m.content {
			o.buf.WriteString(fmt.Sprint(v))
		}
	case msgColor:
		o.buf.WriteString("\x1b[")
		o.buf.WriteString(m.content[0].(string))
		o.buf.WriteRune('m')
	}

	switch m.t {
	case msgRaw, msgPrintln, msgColor:
		// nothing
	default:
		o.buf.WriteRune('\n')
	}

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

func (o *Logger) changeFile(t *time.Time) {

	if t == nil {
		now := time.Now()
		t = &now
	}

	filename := o.fileFunc(t)
	if filename == o.filePrev {
		return
	}

	file, err := openFile(filename, true)
	if err != nil {
		o.err = err
		o.filePrev = filename
		return
	}

	o.file.Sync()
	o.file.Close()

	o.file = file
}
