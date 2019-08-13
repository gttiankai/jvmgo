package heap

import (
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/rtda"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *classfile.ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *rtda.Slot
}

func (self *Class) isSubClassOf(class *Class) bool {

}

func (self *Class) getPackageName() string {

}
