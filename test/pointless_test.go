package j_test

import (
	"fmt"
	"testing"

	"github.com/zhengkai/j"
)

func testPointless(t *testing.T) {

	c := newCapturer()

	m := j.GetDefault()
	for k, v := range m {
		fmt.Printf("%10s: %v\n", k, v)
	}

	fmt.Printf("%v", j.CallerNone)
	fmt.Printf("%v", j.CallerShort)
	fmt.Printf("%v", j.CallerLong)

	fmt.Printf("%v", j.CallerNone-1)

	c.end()
}
