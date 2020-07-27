package c80machine

import (
	"fmt"
)

func (receiver *Type) Reveal() {
	fmt.Println(receiver.Serialize())
}
