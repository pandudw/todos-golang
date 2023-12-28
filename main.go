package main

import (
	"github.com/labstack/echo"
	"github.com/pandudw/todos-golang/database"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Start(":8080")

}
