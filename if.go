package j

import (
	"bytes"
	"os"
	"time"
)

// Config ...
type Config struct {
	Filename   string
	Echo       bool
	Append     bool
	Time       ConfigTimeFormat
	TimeFormat string
	Tunnel     bool
}

// Logger ...
type Logger struct {
	file       *os.File
	enable     bool
	echo       bool
	time       ConfigTimeFormat
	timeFormat string
	err        error
	buf        *bytes.Buffer
	stop       bool
	useTunnel  bool
	tunnel     chan *msg
}

type msg struct {
	isPrintf bool
	time     *time.Time
	content  []string
}

// Close ...
func (o *Logger) Close() {

}

// Enable ...
func (o *Logger) Enable(is bool) {
	o.enable = is
}
