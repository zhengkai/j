package j

import "errors"

// Errors
var (
	ErrTunnelOverflow = errors.New(`Tunnel overflow, log dropped`)
	ErrFileNameEmpty  = errors.New(`Filename is empty`)
)

// ErrorFunc is callback when trigger error
type ErrorFunc func(o *Logger)

func (o *Logger) triggerError(err error) {

	o.Error = err

	fn := o.errorFn
	if fn != nil {
		fn(o)
	}
}
