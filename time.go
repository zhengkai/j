package j

import "time"

// time format in config
const (
	TimeMS = ConfigTimeFormat(iota)
	TimeCustom
	TimeNone
	TimeDayMS
	TimeMonthMS
	TimeYearMS
	Time
	TimeDay
	TimeMonth
	TimeYear
	TimeNS
)

// ConfigTimeFormat ...
type ConfigTimeFormat uint

func (o *Logger) genTime(t *time.Time) (s string) {

	if t == nil {
		ti := time.Now()
		t = &ti
	}

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
	case TimeNS:
		s = t.Format(`15:04:05.000000000 `)
	}

	return
}
