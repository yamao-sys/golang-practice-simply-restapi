package main

import (
	"context"
	"database/sql"
	"github.com/todos-restapi/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"os"
	"fmt"
)

func main() {
	dsn := "root:" + os.Getenv("MYSQL_ROOT_PASSWORD") + "@tcp(db:3306)/" + os.Getenv("MYSQL_DATABASE") + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	user := &models.User{
		UUID: "123456789",
	}
	user.Insert(context.Background(), db, boil.Infer())
}
