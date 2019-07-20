package classfile

type ConstantPool []ConstantInfo


func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
 }

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo  {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)

	name := self.getutf8(ntInfo.nameIndex)
	_type := self.getutf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getutf8(classInfo.nameIndex)
}

func (self ConstantPool) getutf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

