package j

import (
	"bytes"
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

	opClose = opType(iota + 1)
	opEnable
	opDisable
)

// Config ...
type Config struct {
	Filename string
	Echo     bool // stdout
	Append   bool
	Prefix   string
	Time     string // format if Time == TimeCustom
	Tunnel   int    // channel buffer size
}

// Logger ...
type Logger struct {
	file       *os.File
	enable     bool
	echo       bool
	useTime    bool
	timeFormat string
	usePrefix  bool
	prefix     string
	err        error
	buf        *bytes.Buffer
	stop       bool
	stopWait   *sync.WaitGroup
	useTunnel  bool
	tunnel     chan *msg
}

type msg struct {
	t       msgType
	op      opType
	raw     bool
	time    *time.Time
	content []interface{}
	stop    bool
}

type msgType uint8
type opType uint8

// Close ...
func (o *Logger) Close() {
	if o.stop {
		return
	}
	o.stop = true

	if !o.useTunnel {
		return
	}

	w := &sync.WaitGroup{}
	o.stopWait = w
	w.Add(1)
	o.tunnel <- &msg{
		op: opClose,
	}
	w.Wait()

	if o.file != nil {
		o.file.Sync()
		o.file.Close()
		o.file = nil
	}
}

// Enable ...
func (o *Logger) Enable(is bool) {
	if o.stop {
		return
	}

	if !o.useTunnel {
		o.enable = is
		return
	}

	op := opEnable
	if !is {
		op = opDisable
	}

	o.tunnel <- &msg{
		op: op,
	}
}
