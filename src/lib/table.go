package goshd

type Table struct {
	BasicDefinition
	BasicRelation
}

func (self *Table) Init(name string) *Table {
	self.BasicDefinition.Init(name)
	self.BasicRelation.Init()
	return self
}
