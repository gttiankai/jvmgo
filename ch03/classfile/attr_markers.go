package classfile

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// do nothing
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticSttribute struct {
	MarkerAttribute
}
