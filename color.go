package j

// Color Set ANSI color
func (o *Logger) Color(s string) (err error) {
	return o.sendLog(msgColor, s)
}

// ColorOnce Set ANSI color, only next one log
func (o *Logger) ColorOnce(s string) (err error) {
	return o.sendLog(msgColorOnce, s)
}

// ColorReset clean ANSI color
func (o *Logger) ColorReset() (err error) {
	return o.sendLog(msgColor, `0`)
}
