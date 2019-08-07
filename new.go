package j

import (
	"bytes"
	"time"
)

// Config ...
type Config struct {
	Filename string
	Echo     bool // stdout
	Append   bool
	Prefix   string
	Time     string // format if Time == TimeCustom
	Tunnel   int    // channel buffer size
	FileFunc func(t *time.Time) string
}

// New create a new logger with filename
func New(filename string) (o *Logger, err error) {
	config := &Config{
		Filename: filename,
		Time:     TimeMS,
	}
	return NewCustom(config)
}

// NewFunc create a new logger with FileFunc
func NewFunc(fn func(t *time.Time) string) (o *Logger, err error) {
	config := &Config{
		FileFunc: fn,
		Append:   true,
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
