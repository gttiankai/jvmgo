package rtda

// stack frame
type Frame struct {
	lower			*Frame
	localVars 		LocalVars
	operandStack	*OpenrandStack
}

func newFrame(maxLocals , maxStack uint) *Frame {
	return &Frame{
		localVars:    ,
		operandStack: ,
	}
}
