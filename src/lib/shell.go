package goshd

import (
	"context"
	"github.com/jackc/pgx/v4"
	"io"
	"strings"
)

type Shell struct {
	db *pgx.Conn
	context context.Context
	in io.Reader
	out io.Writer
	user *User
	path []*Dir
}

func (self *Shell) Init(db *pgx.Conn, in io.Reader, out io.Writer) *Shell {
	self.db = db
	self.context = context.Background()
	self.in = in
	self.out = out
	return self
}

func (self *Shell) Login(name string) error {
	var u User
	
	row := self.db.QueryRow(self.context,
		"SELECT Name, LastLogin, LastPath, PublicKey FROM Users WHERE Name=$1", name)
	
	if err := u.Load(row); err != nil {
		return err
	}
	
	self.user = &u
	return nil
}

func FormatPath(path []*Dir) string {
	var out strings.Builder
	
	for _, d := range path {
		out.WriteString(d.Name)
		out.WriteRune('>')
	}

	return out.String()
}

func (self *Shell) FormatPath() string {
	return FormatPath(self.path)
}
