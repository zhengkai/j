package main

import (
	"fmt"
	"j"
	"time"
)

func testColor(o *j.Logger) {
	o.BR()

	colorList := []string{
		`1;93;100`,
		`38;2;200;255;100`,
		`38;2;100;255;200`,
		`38;2;255;200;100`,
		`38;2;255;100;200`,
		`38;2;200;100;255`,
		`38;2;100;200;255`,
	}
	for _, v := range colorList {

		s := `` + v
		o.Color(s)
		o.Log(`color`, s)
		o.ColorReset()
	}

	o.Close()
	w.Done()
}

func testNum(o *j.Logger) {

	o.BR()
	for i := 0; i < 100; i++ {
		o.Log(`abc`, fmt.Sprintf(`x: %d`, i))
	}
	o.Log(`end`)

	o.Close()
	w.Done()
}

func testFly(o *j.Logger, color string) {

	useColor := len(color) > 0

	o.BR()
	i := 0
	for {
		i++

		if useColor && i%3 == 1 {
			o.ColorOnce(color)
		}
		o.Log(randStr(20), i)
		time.Sleep(time.Second)

		if i > 10000000000000 {
			break
		}
	}
	o.Log(`end`)

	o.Close()
	w.Done()
}
