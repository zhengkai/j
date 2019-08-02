package main

import (
	"fmt"
	"j"
	"time"
)

func main() {
	x, err := j.New(&j.Config{
		Filename: `test`,
		Echo:     true,
		Time:     j.TimeNS,
		Tunnel:   true,
	})

	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		x.Log(`abc`, fmt.Sprintf(`x: %d`, i))
	}
	x.Log(`end`)

	time.Sleep(10 * time.Second)
}
