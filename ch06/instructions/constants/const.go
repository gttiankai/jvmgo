package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 0x01:push null
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 0x02:push int constant
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// 0X09: push long
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

// 0X0B: push float
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 0X0E push double
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}
