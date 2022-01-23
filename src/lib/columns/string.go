package columns

import (
	"github.com/codr7/goshd/src/lib"
)

var String StringType

func init() {
	String.Init("String")
}

type StringType struct {
	goshd.BasicColumnType
}

func (self *StringType) InitField() interface{} {
	return new(*string)
}
