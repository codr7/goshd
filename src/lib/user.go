package goshd

import (
	"github.com/jackc/pgx/v4"
	"time"
)

type User struct {
	BasicModel
	Name string
	LastLoginAt time.Time
	LastPath []*Dir
	PublicKey []byte
}

func (self *User) Load(src pgx.Row) error {
	self.exists = true
	return src.Scan(&self.Name, &self.LastLoginAt, &self.LastPath, &self.PublicKey)
}
