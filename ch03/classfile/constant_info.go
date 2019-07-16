package classfile

const (
	CONSTANT_Class					= 7
	CONSTANT_Fieldref				= 9
	CONSTANT_Methodref				= 10
	CONSTANT_InterfaceMethodref		= 11
	CONSTANT_String					= 8
	CONSTANT_Integer				= 3
	CONSTANT_Float					= 4
	CONSTANT_Long					= 5
	CONSTANT_Double					= 6
	CONSTANT_NameAndType			= 12
	CONSTANT_Utf8					= 1
	CONSTANT_MethodHandle			= 15
	CONSTANT_MethodType				= 16
	CONSTANT_InvokeDynamic			= 18
)

/*
cp_info {

	u1 tag;
	u1 info[];
}
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := new
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {

	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{}
	}

}

