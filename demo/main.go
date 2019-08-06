package main

import (
	"j"
	"sync"
)

var (
	w *sync.WaitGroup
)

func main() {

	w = &sync.WaitGroup{}

	x, _ := j.NewCustom(&j.Config{
		Filename: `test-x`,
		Time:     j.TimeNS,
		Tunnel:   1000,
	})

	y, _ := j.NewCustom(&j.Config{
		Filename: `test-y`,
		Time:     j.TimeNS,
		Tunnel:   0,
	})

	z1, _ := j.NewCustom(&j.Config{
		Filename: `test-z1`,
		Prefix:   `[prefix]`,
		Time:     j.TimeNS,
		Tunnel:   0,
	})

	z2, _ := j.NewCustom(&j.Config{
		Filename: `test-z2`,
		Prefix:   `[prefix]`,
		Tunnel:   0,
	})

	w.Add(4)
	go testNum(x)
	go testNum(y)
	go testNum(z1)
	go testNum(z2)

	c, _ := j.NewCustom(&j.Config{
		Filename: `test-color`,
		Time:     j.TimeNS,
		Tunnel:   1000,
		Append:   true,
	})
	w.Add(1)
	go testColor(c)

	w.Wait()
}
