package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

/*
ConstantFiledrefInfo 继承自 ConstanMemberrefInfo
*/
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

/*
ConstantMethodrefInfo 继承自 ConstanMemberrefInfo
*/
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

/*
ConstantInterfaceMethodrefInfo 继承自 ConstanMemberrefInfo
*/
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
