package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// branch if int comparison with zero succeeds
// ifeq
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute (frame *rtda.Frame)  {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

// ifne
type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *rtda.Frame)  {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

// iflt
type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame)  {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}
// ifle
type IFLE struct  {
	base.BranchInstruction
}
//
func (self *IFLE) Execute(frame *rtda.Frame)  {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

// ifgt
type IFGT struct {
	base.BranchInstruction
}
func (self *IFGT) Execute(frame *rtda.Frame) {

	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

// ifge
type IFGE struct {

	base.BranchInstruction
}
func (self *IFGE) Execute(frame *rtda.Frame) {

	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
