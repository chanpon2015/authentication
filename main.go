package main

import (
	"fmt"

	"github.com/chanpon2015/authentication/infra"
	"github.com/chanpon2015/authentication/model"
	"github.com/chanpon2015/authentication/usecase/auth"
	jwt "github.com/dgrijalva/jwt-go"
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
		return
	}
	defer db.Close()
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
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"test": "test",
		},
	)
	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		fmt.Println(err)
		return
	}
	token1, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println(ok)
		}
		return []byte("test"), nil
	})
	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["test"])
	}
	fmt.Println(tokenString)
}
