package j

import "os"

var (
	flagNew    = os.O_RDWR | os.O_CREATE
	flagAppend = os.O_APPEND | os.O_CREATE | os.O_WRONLY
)

// New create a new logger
func New(c *Config) (o *Logger, err error) {

	o = &Logger{
		enable: true,
	}

	if len(c.Filename) > 0 {

		flag := flagNew
		if c.Append {
			flag = flagAppend
		}

		o.file, err = os.OpenFile(c.Filename, flag, 0644)
		if err != nil {
			return
		}
	}

	return
}
