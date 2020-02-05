package zj_test

import (
	"regexp"
	"testing"

	"github.com/zhengkai/zj"
)

func testLineFunc(t *testing.T) {

	re := regexp.MustCompile(`(foo|bar)`)
	repl := " > $1 < "
	fn := func(s *string) {
		r := re.ReplaceAllString(*s, repl)
		*s = r
	}

	c := newCapturer()

	x := zj.New(&zj.Config{
		LineFn: fn,
	})
	x.Log(`123fo`, `obar321`)
	x.Print(`123fo`, `obar321`)
	x.Compact(`123fo`, `obar321`)
	x.Raw(`123foobar321`)

	s := c.end()
	x.Close()

	replaceTime(&s)
	replaceCaller(&s)

	str := `[TIME] [CALLER] 123fo o > bar < 321
[TIME] [CALLER] 123 > foo <  > bar < 321
[TIME] [CALLER] 123 > foo <  > bar < 321
123foobar321`

	if s != str {
		t.Error(`callback LineFunc fail`)
	}
}
