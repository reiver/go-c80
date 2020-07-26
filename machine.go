package c80

import (
	"github.com/reiver/go-dast8x8"

	"github.com/reiver/go-c80/machine"
)

var (
	machine c80machine.Type
)

func init() {
	if err := Colorize("vt"); nil != err {
		panic(err)
	}

	{
		const diff = 247

		for id:=uint8(0); id<2; id++ {
			localID := id+diff

			if err := SetSprite("8x8", localID, dast8x8.Sprite(id)); nil != err {
				panic(err)
			}
		}
	}
}
