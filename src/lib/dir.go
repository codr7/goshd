package goshd

import (
	"time"
)

type Dir struct {
	Id uint64
	Name string
	Version uint64
	CreatedAt time.Time
	ModifiedAt time.Time
}
