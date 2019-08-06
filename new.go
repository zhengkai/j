package j

import (
	"bytes"
	"os"
)

const (
	flagNew    = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	flagAppend = os.O_CREATE | os.O_WRONLY | os.O_APPEND
)

// New create a new logger with filename
func New(filename string) (o *Logger, err error) {
	config := &Config{
		Filename: filename,
		Time:     TimeMS,
	}
	return NewCustom(config)
}

// NewCustom create a new logger with config
func NewCustom(c *Config) (o *Logger, err error) {

	o = &Logger{
		enable:    true,
		echo:      c.Echo,
		buf:       &bytes.Buffer{},
		useTunnel: c.Tunnel > 0,
	}

	if len(c.Filename) > 0 {
		o.file, err = openFile(c.Filename, c.Append)
		if err != nil {
			o = nil
			return
		}
	}

	if len(c.Time) > 0 {
		o.useTime = true
		o.timeFormat = c.Time
	}

	if len(c.Prefix) > 0 {
		o.usePrefix = true
		o.prefix = c.Prefix
	}

	if o.useTunnel {
		o.tunnel = make(chan *msg, c.Tunnel)
		go o.bgLog()
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
