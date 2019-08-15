package heap

import (
	"jvmgo/ch06/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp 			*classpath.ClassPath
	classMap 	map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {

	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class  {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}
	return self.
}

func (self *ClassLoader) loadNonArrayClass(name string)  {

}
