package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	readInfo(read *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameindex := reader.readUint16()
	attrName := cp.getutf8(attrNameindex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code" : return
	case "ConstantValue"
	case "Deprecated"
	case "Exception"
	case "LineNumberTable"
	case "LocalVariableTable"
	case "SourceFuke"
	case "Synthetic"
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

