package j

// Color Set ANSI color
func (o *Logger) Color(s string) (err error) {
	return o.sendLog(msgColor, s)
}

// ColorReset clean ANSI color
func (o *Logger) ColorReset() (err error) {
	return o.sendLog(msgColor, `0`)
}
