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
	o = New(config)
	return
}

// NewFile create a new logger with filename
func NewFile(filename string) (o *Logger) {
	config := &Config{
		Filename: filename,
	}
	return New(config)
}

// NewFunc create a new logger with FileFunc
func NewFunc(fileFn FileFunc) (o *Logger) {
	config := &Config{
		FileFn: fileFn,
		Append: true,
	}
	return New(config)
}

// New create a new logger
func New(c *Config) (o *Logger) {

	applyConfig(c)

	return NewPure(c)
}

// NewPure create a new logger without default config
func NewPure(c *Config) (o *Logger) {

	applyFileConfig(c)

	o = &Logger{
		enable:    true,
		echo:      c.Echo,
		buf:       &bytes.Buffer{},
		useTunnel: c.Tunnel > 0,
		lineFn:    c.LineFn,
		caller:    c.Caller,
		permFile:  c.PermFile,
		permDir:   c.PermDir,
		errorFn:   c.ErrorFn,
	}

	if c.File == nil {

		o.fileSelf = true

		if c.FileFn != nil {
			o.fileFn = c.FileFn
			now := time.Now()
			c.Filename = c.FileFn(&now)
			c.Append = true
		}
		if len(c.Filename) > 0 {
			var err error
			o.file, err = o.openFile(c.Filename, c.Append)
			if err != nil {
				o.stop = true
				o.triggerError(err)
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
