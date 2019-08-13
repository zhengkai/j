package j_test

import (
	"regexp"
	"time"

	"github.com/zhengkai/j"
)

func ExampleNewFunc() {
	jx := j.NewFunc(func(t *time.Time) string {
		return t.Format(`log/01-02/15.log`)
	})
	jx.Log(`file func`)
}

func ExampleLineFunc() {
	re := regexp.MustCompile(`(foo|bar)`)
	repl := " [ $1 ] "
	fn := func(s *string) {
		r := re.ReplaceAllString(*s, repl)
		*s = r
	}

	x := j.New(&j.Config{
		LineFn: fn,
	})

	x.Log(`afoob`)
	// output:
	// 17:04:31.829 example_test.go:29 a [ foo ] b
}
