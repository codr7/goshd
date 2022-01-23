package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v4"
	"github.com/codr7/goshd/src/lib"
	"github.com/codr7/goshd/src/lib/columns"
	"github.com/codr7/goshd/src/lib/db"
)

func main() {
	url := "postgres://goshd:goshd@localhost:5432/goshd"
	dbc, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Fatal(err)
	}

	defer dbc.Close(context.Background())

	db.Users.Init("Users")
	db.UserName.Init("Name", &columns.String)
	db.Users.Append(&db.UserName)
	//godhd.UserLastLoginAt.Init("LastLoginAt")
	//goshd.Users.Append(&goshd.UserLastLoginAt)
	//godhd.UserLastPath.Init("LastPath")
	//goshd.Users.Append(&goshd.UserLastPath)
	
	var s goshd.Shell
	if err := s.Init(dbc, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}

	if err := s.Login("admin"); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%v\n", s.FormatPath())
	
	var u goshd.User
	u.Name = "admin"
}
