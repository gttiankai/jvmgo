package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {

	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

// branch if int comparison succeeds
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {

	if val1, val2 := _icmpPop(frame); val1 == val2 {

		base.Branch(frame, self.Offset)
	}
}

// if_icmpne
type IF_ICMPNE struct {
	base.BranchInstruction
}

// Execute ...
func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

// if_icmplt
type IF_ICMPLT struct {
	base.BranchInstruction
}

// Execte ...
func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

// if_icmple
type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {

	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

// if_icmpgt
type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

// if_icmpge
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {

	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}
