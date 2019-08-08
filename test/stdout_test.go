package j_test

import (
	"testing"

	"github.com/zhengkai/j"
)

func TestEcho(t *testing.T) {

	x := j.NewEcho()
	x.Log(`abc`)

	// t.Fatal("not implemented")
}
