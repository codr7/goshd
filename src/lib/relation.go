package goshd

type Relation interface {
	Columns() []*Column
}

type BasicRelation struct {
	columns []*Column
	columnLookup map[*Column]int
}

func (self *BasicRelation) Init() *BasicRelation {
	self.columnLookup = make(map[*Column]int)
	return self
}

func (self *BasicRelation) Append(cols...*Column) {
	for i := 0; i < len(cols); i++ {
		self.columnLookup[cols[i]] = len(self.columns)+i
	}
	
	self.columns = append(self.columns, cols...)
}

func (self *BasicRelation) IndexOf(col *Column) int {
	return self.columnLookup[col]
}

func (self *BasicRelation) NewRecord() []interface{} {
	out := make([]interface{}, len(self.columns))

	for i, c := range self.columns {
		out[i] = c.Type().InitField()
	}

	return out
}
