package goshd

import (
	"github.com/jackc/pgx/v4"
)

type Model interface {
	Exists() bool
	Load(src pgx.Row) error
}

type BasicModel struct {
	exists bool
}

func (self *BasicModel) Exists() bool {
	return self.exists
}
