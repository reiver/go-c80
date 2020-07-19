package c80

import (
	"github.com/reiver/go-rgba32"

	"github.com/reiver/go-c80/machine"
)

var (
	machine c80machine.Type
)

func init() {
	if err := Colorize("vt"); nil != err {
		panic(err)
	}
}
