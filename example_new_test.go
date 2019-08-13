package j_test

import (
	"time"

	"github.com/zhengkai/j"
)

func ExampleNewFunc() {

	jx := j.NewFunc(func(t *time.Time) string {
		return t.Format(`log/01-02/15.log`)
	})
	jx.Log(`file func`)
}
