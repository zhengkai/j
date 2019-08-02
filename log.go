package j

import (
	"fmt"
	"time"
)

// Log ...
func (o *Logger) Log(content string) {

	if !o.enable {
		return
	}

	var s string

	if o.time != TimeNone {

		t := time.Now()

		switch o.time {
		case TimeMS:
			s = t.Format(`15:04:05.000 `)
		case TimeDayMS:
			s = t.Format(`02 15:04:05.000 `)
		case TimeMonthMS:
			s = t.Format(`01-02 15:04:05.000 `)
		case TimeYearMS:
			s = t.Format(`2006-01-02 15:04:05.000 `)
		case TimeDay:
			s = t.Format(`02 15:04:05 `)
		case TimeMonth:
			s = t.Format(`01-02 15:04:05 `)
		case TimeYear:
			s = t.Format(`2006-01-02 15:04:05 `)
		}
	}

	if o.file != nil {
		o.file.WriteString(s)
	}

	if o.echo {
		fmt.Println(s)
	}
}

// Enable ...
func (o *Logger) Enable(is bool) {
	o.enable = is
}
