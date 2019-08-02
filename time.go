package j

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
)

// ConfigTimeFormat ...
type ConfigTimeFormat uint
