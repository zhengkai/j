package j

// Color Set ANSI color
func (o *Logger) Color(s string) (err error) {
	return o.sendLog(msgColor, s)
}

// ColorOnce Set ANSI color, only next line
func (o *Logger) ColorOnce(s string) (err error) {
	return o.sendLog(msgColorOnce, s)
}

// ColorStop clean ANSI color
func (o *Logger) ColorStop() (err error) {
	return o.sendLog(msgColor, `0`)
}
