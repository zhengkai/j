package zj

import "sync"

// Close permanently disable Log() series.
// log files will be closed, unless come from SetFile() or config File.
//
// Note: if create by config Tunnel, must use Close() to
// wait that all logs have been processed
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
		if o.fileSelf {
			o.file.Close()
		}
		o.file = nil
	}
}

// Enable used to temporarily disable
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
