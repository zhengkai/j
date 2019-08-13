package j

// Color use ANSI color for each line
func (o *Logger) Color(s string) (err error) {
	return o.sendLog(msgColor, s)
}

// ColorOnce use ANSI color, only next line
func (o *Logger) ColorOnce(s string) (err error) {
	return o.sendLog(msgColorOnce, s)
}

// ColorStop no longer use ANSI color
func (o *Logger) ColorStop() (err error) {
	return o.sendLog(msgColor, `0`)
}
