package classfile

/*
CONSTANT_MethodHandler_info {
	u1 tag;
	u2 reference_kind;
	u2 reference_index;
}
*/

type ConstantMethodHandleInfo struct {
	referenceKind  uint16
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint16()
	self.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_Info {
	u1 tag
	u2 descriptor_index
}
*/

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
	u1 tag;
	u2 bootstrap_method_attr_index;
	u2 name_and_type
}
*/

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}