package zj_test

import (
	"fmt"
	"testing"

	"github.com/zhengkai/zj"
)

func testPointless(t *testing.T) {

	c := newCapturer()

	m := zj.GetDefault()
	for k, v := range m {
		fmt.Printf("%10s: %v\n", k, v)
	}

	fmt.Printf("%v", zj.CallerNone)
	fmt.Printf("%v", zj.CallerShort)
	fmt.Printf("%v", zj.CallerLong)

	fmt.Printf("%v", zj.CallerNone-1)

	c.end()
}
