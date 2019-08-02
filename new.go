package j

import (
	"bytes"
	"fmt"
	"os"
)

const (
	flagNew    = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	flagAppend = os.O_CREATE | os.O_WRONLY | os.O_APPEND
)

// New create a new logger
func New(c *Config) (o *Logger, err error) {

	o = &Logger{
		enable: true,
		// echo:   c.Echo,
		buf:       &bytes.Buffer{},
		time:      c.Time,
		useTunnel: c.Tunnel,
	}

	if len(c.Filename) > 0 {
		o.file, err = openFile(c.Filename, c.Append)
		if err != nil {
			o = nil
			return
		}
	}

	if c.Time == TimeCustom {
		o.timeFormat = c.TimeFormat
	}

	if o.useTunnel {
		o.tunnel = make(chan *msg, 1000)
		fmt.Println(`start bgLog`)
		go func() {
			o.bgLog()
			fmt.Println(`start bgLog end`)
		}()
	}

	return
}

func openFile(filename string, isAppend bool) (f *os.File, err error) {

	flag := flagNew
	if isAppend {
		flag = flagAppend
	}

	f, err = os.OpenFile(filename, flag, 0644)
	if err != nil {
		return
	}

	return
}
