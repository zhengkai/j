package j_test

import (
	"time"

	"github.com/zhengkai/j"
)

func ExampleNewFunc() {

	jx, _ := j.NewFunc(func(t *time.Time) string {
		return t.Format(`log/same-01-02/15/04`)
	})
	jx.Log(`file func`)
}
