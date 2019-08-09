package j

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	msgPrintln = msgType(iota)
	msgPrintf
	msgPrint
	msgCompact
	msgRaw
	msgColor
	msgColorOnce

	opClose = opType(iota + 1)
	opEnable
	opDisable
)

// just like Lshortfile/Llongfile in pkg/log
const (
	CallerNone = callerType(iota)
	CallerShort
	CallerLong
)

type msgType uint8
type opType uint8
type callerType uint8

// Logger ...
type Logger struct {
	Error error

	file       *os.File
	fileSelf   bool
	enable     bool
	echo       bool
	useTime    bool
	timeFormat string
	usePrefix  bool
	prefix     string
	buf        *bytes.Buffer
	stop       bool
	stopWait   *sync.WaitGroup
	useColor   bool
	stopColor  bool
	color      string
	colorOnce  bool
	useTunnel  bool
	tunnel     chan *msg
	fileFunc   func(t *time.Time) (filename string)
	filePrev   string
	lineFunc   func(line *string)
	caller     callerType
}

type msg struct {
	t       msgType
	op      opType
	raw     bool
	time    *time.Time
	content []interface{}
	stop    bool
	caller  *caller
}

type caller struct {
	file string
	line int
}

// Log just like log.Println
func (o *Logger) Log(a ...interface{}) (err error) {
	return o.sendLog(msgPrintln, a...)
}

// Logf just like log.Printf
func (o *Logger) Logf(format string, a ...interface{}) (err error) {
	return o.sendLog(msgPrintf, fmt.Sprintf(format, a...))
}

// Print just like log.Print, added linebreak
func (o *Logger) Print(a ...interface{}) (err error) {
	return o.sendLog(msgPrint, a...)
}

// Compact just like fmt.Print, but no spaces
func (o *Logger) Compact(a ...interface{}) (err error) {
	return o.sendLog(msgCompact, a...)
}

// Raw log raw (no time, no linebreak, no spaces)
func (o *Logger) Raw(a ...interface{}) (err error) {
	return o.sendLog(msgRaw, a...)
}

// BR add an empty line
func (o *Logger) BR() (err error) {
	return o.sendLog(msgRaw, "\n")
}

func (c callerType) String() string {

	switch c {

	case CallerNone:
		return `none`
	case CallerShort:
		return `short`
	case CallerLong:
		return `long`
	}

	return `unknown`
}
