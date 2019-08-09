package j

import (
	"bytes"
	"time"
)

// NewEcho create a new logger without file, only stdout
func NewEcho() (o *Logger) {
	config := &Config{
		Echo: true,
	}
	o, _ = New(config)
	return
}

// NewFile create a new logger with filename
func NewFile(filename string) (o *Logger, err error) {
	config := &Config{
		Filename: filename,
	}
	return New(config)
}

// NewFunc create a new logger with FileFunc
func NewFunc(fn func(t *time.Time) (filename string)) (o *Logger, err error) {
	config := &Config{
		FileFunc: fn,
		Append:   true,
	}
	return New(config)
}

// New create a new logger
func New(c *Config) (o *Logger, err error) {

	applyConfig(c)

	return NewPure(c)
}

// NewPure create a new logger without default config
func NewPure(c *Config) (o *Logger, err error) {

	o = &Logger{
		enable:    true,
		echo:      c.Echo,
		buf:       &bytes.Buffer{},
		useTunnel: c.Tunnel > 0,
		lineFunc:  c.LineFunc,
		caller:    c.Caller,
	}

	if c.File == nil {

		o.fileSelf = true

		if c.FileFunc != nil {
			o.fileFunc = c.FileFunc
			now := time.Now()
			c.Filename = c.FileFunc(&now)
		}
		if len(c.Filename) > 0 {
			o.file, err = openFile(c.Filename, c.Append)
			if err != nil {
				o = nil
				return
			}
		}

	} else {

		o.file = c.File
	}

	if len(c.TimeFormat) > 0 {
		o.useTime = true
		o.timeFormat = c.TimeFormat
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
