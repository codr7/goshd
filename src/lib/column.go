package goshd

type ColumnType interface {
	InitField() interface{}
}

type BasicColumnType struct {
	name string
}

func (self *BasicColumnType) Init(name string) {
	self.name = name
}

type Column struct {
	BasicDefinition
	columnType ColumnType
}

func (self *Column) Init(name string, columnType ColumnType) *Column {
	self.BasicDefinition.Init(name)
	self.columnType = columnType
	return self
}

func (self *Column) Type() ColumnType {
	return self.columnType
}



