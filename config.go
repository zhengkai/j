package j

import (
	"os"
)

var (
	configDefault = make(map[configKey]interface{})
)

// Config key, used by SetDefault()
const (
	Echo = configKey(iota + 1)
	Append
	Prefix
	TimeFormat
	Tunnel
	LineFn
	Caller
	PermFile
	PermDir
	ErrorFn
)

// Config for create logger
type Config struct {

	// fill only one of them
	File     *os.File
	FileFn   FileFunc
	Filename string

	Echo       bool // stdout
	Append     bool
	Prefix     string
	TimeFormat string
	Tunnel     int // channel buffer size, see also Close()
	LineFn     LineFunc
	ErrorFn    ErrorFunc
	Caller     callerType
	PermFile   os.FileMode
	PermDir    os.FileMode
}

// ConfigKey for SetDefault
type configKey uint8

func initColor() {
	configDefault[Echo] = true
	configDefault[TimeFormat] = TimeMS
	configDefault[Caller] = CallerShort
	configDefault[PermDir] = os.FileMode(0755)
	configDefault[PermFile] = os.FileMode(0644)
}

func (c configKey) String() string {

	switch c {

	case Echo:
		return `Echo`
	case Append:
		return `Append`
	case Prefix:
		return `Prefix`
	case TimeFormat:
		return `TimeFormat`
	case Tunnel:
		return `Tunnel`
	case LineFn:
		return `LineFn`
	case Caller:
		return `Caller`
	case PermFile:
		return `PermFile`
	case PermDir:
		return `PermDir`
	case ErrorFn:
		return `ErrorFunc`
	}

	return `unknown`
}

// SetDefault for new logger
func SetDefault(k configKey, v interface{}) (ok bool) {

	switch k {

	case Echo, Append:
		var r bool
		if r, ok = v.(bool); ok {
			configDefault[k] = r
		}
	case Prefix, TimeFormat:
		var r string
		if r, ok = v.(string); ok {
			configDefault[k] = r
		}

	case Tunnel:
		var r int
		if r, ok = v.(int); ok {
			configDefault[k] = r
		}

	case LineFn:
		var r LineFunc
		if r, ok = v.(func(*string)); ok {
			configDefault[k] = r
		}

	case Caller:
		var r callerType
		if r, ok = v.(callerType); ok {
			configDefault[k] = r
		}

	case PermDir, PermFile:
		var i int
		var r os.FileMode
		if i, ok = v.(int); ok {
			configDefault[k] = os.FileMode(i)
			return
		}
		if r, ok = v.(os.FileMode); ok {
			configDefault[k] = r
		}

	case ErrorFn:
		var r ErrorFunc
		if r, ok = v.(func(*Logger)); ok {
			configDefault[k] = r
		}
	}

	return
}

// UnsetDefault ...
func UnsetDefault(k configKey) {
	delete(configDefault, k)
}

// GetDefault ...
func GetDefault() (m map[string]interface{}) {

	m = make(map[string]interface{})
	for k, v := range configDefault {
		m[k.String()] = v
	}
	return
}

func applyConfig(c *Config) {

	if !c.Echo {
		v, ok := configDefault[Echo]
		if ok {
			c.Echo = v.(bool)
		}
	}

	if !c.Append {
		v, ok := configDefault[Append]
		if ok {
			c.Append = v.(bool)
		}
	}

	if len(c.Prefix) == 0 {
		v, ok := configDefault[Prefix]
		if ok {
			c.Prefix = v.(string)
		}
	}

	if len(c.TimeFormat) == 0 {
		v, ok := configDefault[TimeFormat]
		if ok {
			c.TimeFormat = v.(string)
		}
	}

	if c.Tunnel == 0 {
		v, ok := configDefault[Tunnel]
		if ok {
			c.Tunnel = v.(int)
		}
	}

	if c.Caller == 0 {
		v, ok := configDefault[Caller]
		if ok {
			c.Caller = v.(callerType)
		}
	}

	if c.LineFn == nil {
		v, ok := configDefault[LineFn]
		if ok {
			c.LineFn = v.(LineFunc)
		}
	}

	if c.ErrorFn == nil {
		v, ok := configDefault[ErrorFn]
		if ok {
			c.ErrorFn = v.(ErrorFunc)
		}
	}
}

func applyFileConfig(c *Config) {

	if c.PermFile == 0 {
		v, ok := configDefault[PermFile]
		if ok {
			c.PermFile = v.(os.FileMode)
		}
	}

	if c.PermDir == 0 {
		v, ok := configDefault[PermDir]
		if ok {
			c.PermDir = v.(os.FileMode)
		}
	}
}
