package main

import (
	"fmt"
	"j"
	"time"
)

func task() {

	taskDir()
	taskColor()
	taskN()
}

func taskColor() {

	c, _ := j.NewCustom(&j.Config{
		Filename: `log/color/text`,
		Time:     j.TimeNS,
		Tunnel:   1000,
		Append:   true,
	})
	w.Add(1)
	go testColor(c)
}

func taskDir() {
	check, err := j.NewCustom(&j.Config{
		Filename: `test-x`,
		FileFunc: func(t *time.Time) string {
			return t.Format(`log/01-02/15/04`)
		},
		Time:   j.TimeNS,
		Tunnel: 1000,
	})

	fmt.Println(`check`, err)
	if err != nil {
		return
	}

	w.Add(1)
	go testFly(check)
}

func taskN() {

	x, _ := j.NewCustom(&j.Config{
		Filename: `test-x`,
		FileFunc: func(t *time.Time) string {
			return t.Format(`log/test/2006-01-02`)
		},
		Time:   j.TimeNS,
		Tunnel: 1000,
	})

	y, _ := j.NewCustom(&j.Config{
		Filename: `test-y`,
		Time:     j.TimeNS,
		Tunnel:   0,
	})

	z1, _ := j.NewCustom(&j.Config{
		Filename: `log/test-z1`,
		Prefix:   `[prefix] `,
		Time:     j.TimeNS,
		Tunnel:   0,
	})

	z2, _ := j.NewCustom(&j.Config{
		Filename: `log/test-z2`,
		Prefix:   `[prefix] `,
		Tunnel:   0,
	})

	w.Add(5)
	go testNum(x)
	go testNum(y)
	go testNum(z1)
	go testNum(z2)
}

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

func testFly(o *j.Logger) {

	o.BR()
	i := 0
	for {
		i++
		o.Log(randStr(20), i)
		time.Sleep(time.Second / 10)

		if i > 10000000000000 {
			break
		}
	}
	o.Log(`end`)

	o.Close()
	w.Done()
}
