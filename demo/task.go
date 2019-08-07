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
	taskSameFile()
	taskColorSameFile()
}

func taskColorSameFile() {

	x1, err := j.NewCustom(&j.Config{
		Filename: `log/color-same`,
		Prefix:   `[ X1 ]`,
		Time:     j.TimeNS,
		Tunnel:   1000,
	})

	if err != nil {
		fmt.Println(`taskSameFile 1`, err)
		return
	}

	/*
		x2, err := j.NewCustom(&j.Config{
			Filename: `log/color-same`,
			Prefix:   `[X2]`,
			Time:     j.TimeMS,
			Tunnel:   1000,
		})

		if err != nil {
			fmt.Println(`taskSameFile 2`, err)
			return
		}
	*/

	w.Add(2)
	go testFly(x1, ``)
	go testFly(x1, `38;2;200;255;100`)
}

func taskSameFile() {

	x1, err := j.NewCustom(&j.Config{
		FileFunc: func(t *time.Time) string {
			return t.Format(`log/same-01-02/15/04`)
		},
		Prefix: `[ X1 ]`,
		Time:   j.TimeNS,
		Tunnel: 1000,
	})

	if err != nil {
		fmt.Println(`taskSameFile 1`, err)
		return
	}

	x2, err := j.NewCustom(&j.Config{
		FileFunc: func(t *time.Time) string {
			return t.Format(`log/same-01-02/15/04`)
		},
		Prefix: `[X2] `,
		Time:   j.TimeMS,
		Tunnel: 1000,
	})

	if err != nil {
		fmt.Println(`taskSameFile 2`, err)
		return
	}

	w.Add(2)
	go testFly(x1, ``)
	go testFly(x2, `38;2;200;255;100`)
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
	go testFly(check, ``)
}

func taskN() {

	x, _ := j.NewCustom(&j.Config{
		FileFunc: func(t *time.Time) string {
			return t.Format(`log/test/2006-01-02`)
		},
		Time:   j.TimeNS,
		Tunnel: 1000,
	})

	y, _ := j.NewCustom(&j.Config{
		Filename: `log-y`,
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
