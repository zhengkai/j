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
	CallerNone    = callerType(iota + 1)
	CallerShort   // caller.go:42
	CallerShorter // caller:42
	CallerLong    // /dir/caller.go:42
)

type msgType uint8
type opType uint8
type callerType uint8

// Logger will be always returned by New series func, if there is any error, recorded in Error
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
	fileFn     FileFunc
	filePrev   string
	lineFn     LineFunc
	caller     callerType
	errorFn    ErrorFunc

	permFile os.FileMode
	permDir  os.FileMode
}

// FileFunc is the type of the function called for getting filename.
// If the filename is different from the previous one,
// a new log file will be created
type FileFunc func(t *time.Time) (filename string)

// LineFunc is the type of the function called by from
// Log / Logf / Print / Compact, but Raw.
type LineFunc func(line *string)

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

// Raw log raw (no time, no linebreak, no spaces),
// will not trigger lineFn(LineFunc)
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
