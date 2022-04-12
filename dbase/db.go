package dbase

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-redis/redis"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DbSql *sql.DB
	DbRed redis.Cmdable
}

func CheckDB() Database {
	_, err := os.Stat("dbase/database-sqlite.db")
	if os.IsNotExist(err) {
		createFile()
	}
	var d Database
	d.open("dbase/database-sqlite.db")
	d.createTable()
	d.DbRed, _ = CreateRedis()
	return d
}
func CreateRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set("counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	return rdb, nil
}

func createFile() {
	file, err := os.Create("dbase/database-sqlite.db")
	if err != nil {
		log.Fatalf("file doesn't create %v", err)
	}
	defer file.Close()
}

func (d *Database) open(file string) {
	var err error
	d.DbSql, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("this error is in dbase/open() %v", err)
	}
}

func (d *Database) createTable() {
	_, err := d.DbSql.Exec(`CREATE TABLE IF NOT EXISTS user (
        "id"    INTEGER NOT NULL UNIQUE,
        "name"    TEXT NOT NULL UNIQUE,
        "sername"    TEXT NOT NULL,
        PRIMARY KEY("id" AUTOINCREMENT)
    );`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE users")
	}

}
