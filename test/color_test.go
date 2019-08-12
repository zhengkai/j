package j_test

import (
	"regexp"
	"testing"

	"github.com/zhengkai/j"
)

var (
	reTime     = regexp.MustCompile(`\d{2}:\d{2}:\d{2}\.\d{3} `)
	reColor    = regexp.MustCompile(`\x1b\[[0-9;]+m`)
	reColorEnd = regexp.MustCompile(`\x1b\[0m`)
)

func testColor(t *testing.T) {

	filename := `log-color`

	c := newCapturer()
	x := j.NewFile(filename)
	if x.Error != nil {
		t.Error(`method "NewFile" fail`)
	}

	color := `38;2;200;255;100`

	x.Color(color)
	x.Log(`1 color test`)
	x.Log(`2 color test`, `again`)
	x.ColorStop()

	x.Log(`3 no color`)

	x.ColorOnce(color)
	x.Log(4, `color once`)

	x.Raw(`raw`)

	x.Log(`5 no color`)

	s := c.end()

	sr := s

	replaceTime(&sr)
	replaceCaller(&sr)
	replaceColor(&sr)

	str := `[COLOR_START][TIME] [CALLER] 1 color test[COLOR_END]
[COLOR_START][TIME] [CALLER] 2 color test again[COLOR_END]
[TIME] [CALLER] 3 no color
[COLOR_START][TIME] [CALLER] 4 color once[COLOR_END]
raw[TIME] [CALLER] 5 no color` + "\n"

	if sr != str {
		t.Error(`method "Color" fail`, str, sr)
	}

	sf, err := loadFile(filename)
	if err != nil {
		t.Error(`load file fail`, filename, err)
	}

	if s != sf {
		t.Error(`write file not match`, sf)
	}

	sc := `continue`

	x2 := j.New(&j.Config{
		File: x.GetFile(),
	})

	c = newCapturer()
	x2.Raw(sc)
	c.end()

	sf, err = loadFile(filename)
	if err != nil {
		t.Error(`load file fail`, filename, err)
	}

	s += sc
	if s != sf {
		t.Error(`write file not match`, sf)
	}
}
