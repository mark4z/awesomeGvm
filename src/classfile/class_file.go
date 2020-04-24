package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (c *ClassFile) read(reader *ClassReader) {

}

func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

//getter
func (c *ClassFile) MinorVersion() uint16 {

}

//getter
func (c *ClassFile) MajorVersion() uint16 {

}

//getter
func (c *ClassFile) ConstantPool() ConstanPool {

}

//getter
func (c *ClassFile) AccessFlags() uint16 {

}

//getter
func (c *ClassFile) Fields() []*MemberInfo {

}

//getter
func (c *ClassFile) Methods() []*MemberInfo {

}

//getter
func (c *ClassFile) ClassName() string {

}

//getter
func (c *ClassFile) SuperClassName() string {

}

//getter
func (c *ClassFile) InterfaceNames() []string {

}
