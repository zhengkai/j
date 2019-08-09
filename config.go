package j

import (
	"os"
	"time"
)

var (
	configDefault = &ConfigDefault{
		PermDir:  0755,
		PermFile: 0644,
	}
)

// Config for create logger
type Config struct {
	Filename string
	FileFunc func(t *time.Time) (filename string)

	// below same as ConfigDefault
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

// ConfigDefault ...
type ConfigDefault struct {
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

func applyConfig(c *Config) {

	if !c.Echo {
		c.Echo = configDefault.Echo
	}
	if !c.Echo {
		c.Echo = configDefault.Append
	}
	if len(c.Prefix) == 0 {
		c.Prefix = configDefault.Prefix
	}
	if len(c.TimeFormat) == 0 {
		c.TimeFormat = configDefault.TimeFormat
	}
	if c.Tunnel == 0 {
		c.Tunnel = configDefault.Tunnel
	}
	if c.LineFunc == nil {
		c.LineFunc = configDefault.LineFunc
	}
	if c.Caller == 0 {
		c.Caller = configDefault.Caller
	}
	if c.PermFile == 0 {
		c.PermFile = configDefault.PermFile
	}
	if c.PermDir == 0 {
		c.PermDir = configDefault.PermDir
	}
}
