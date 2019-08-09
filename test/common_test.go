package j_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func loadFile(filename string) (s string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	s = string(b)
	return
}

func replaceTime(s *string) {
	out := reTime.ReplaceAllString(*s, `[TIME] `)
	*s = out
}

func replaceCaller(s *string) {

	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	_, file = filepath.Split(file)
	file = strings.Replace(file, `.`, `\.`, 1)

	pattern := file + `:\d+`

	re := regexp.MustCompile(pattern)
	out := re.ReplaceAllString(*s, `[CALLER]`)
	*s = out
}

func replaceColor(s *string) {
	out := reColorEnd.ReplaceAllString(*s, `[COLOR_END]`)
	out = reColor.ReplaceAllString(out, `[COLOR_START]`)
	*s = out
}
