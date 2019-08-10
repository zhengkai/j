package j

import (
	"os"
	"time"
)

var (
	configDefault = make(map[configKey]interface{})
)

// Config key
const (
	Echo = configKey(iota + 1)
	Append
	Prefix
	TimeFormat
	Tunnel
	LineFunc
	Caller
	PermFile
	PermDir
)

// Config for create logger
type Config struct {
	File       *os.File
	Filename   string
	FileFunc   func(t *time.Time) (filename string)
	Echo       bool // stdout
	Append     bool
	Prefix     string
	TimeFormat string
	Tunnel     int // channel buffer size
	LineFunc   func(line *string)
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
	case LineFunc:
		return `LineFunc`
	case Caller:
		return `Caller`
	case PermFile:
		return `PermFile`
	case PermDir:
		return `PermDir`
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

	case LineFunc:
		var r func(line *string)
		if r, ok = v.(func(line *string)); ok {
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

	if c.LineFunc == nil {
		v, ok := configDefault[LineFunc]
		if ok {
			c.LineFunc = v.(func(line *string))
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
