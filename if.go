package j

import "os"

// Config ...
type Config struct {
	Filename   string
	Stdout     bool
	Append     bool
	Time       ConfigTimeFormat
	TimeFormat string
}

// Logger ...
type Logger struct {
	file       *os.File
	enable     bool
	echo       bool
	time       ConfigTimeFormat
	timeFormat string
}
