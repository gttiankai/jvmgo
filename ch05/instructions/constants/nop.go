package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct {
	base.NopInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
