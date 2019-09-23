package main

import (
	"fmt"

	"github.com/chanpon2015/authentication/infra"
	"github.com/chanpon2015/authentication/model"
	"github.com/chanpon2015/authentication/usecase/auth"
)

func main() {
	psql := model.Postgresql{
		Host:     "localhost",
		DbName:   "auth",
		User:     "auth",
		Password: "postgres",
		Port:     5432,
		SSLMode:  "disable",
	}
	db, err := infra.Open(&psql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connect!")
	db.AutoMigrate(&model.User{})
	if err := auth.Add(db, "aaaa", "bbbb", "cccc"); err != nil {
		fmt.Println(err)
	}
	if err := auth.Add(db, "dddd", "bbbb", "eeee"); err != nil {
		fmt.Println(err)
	}
	result := auth.Authentication(db, "aaaa", "bbbb")
	fmt.Println(result)
	result = auth.Authentication(db, "dddd", "bbbb")
	fmt.Println(result)
	defer db.Close()
}
