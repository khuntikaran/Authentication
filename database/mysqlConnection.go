package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "username"
	dbPass := "password"
	dbName := "Auth01"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	ctx, calcleFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer calcleFunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error while connecting to MYSQL database")
	}

	log.Printf("Connected successfully ")
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	return db
}
