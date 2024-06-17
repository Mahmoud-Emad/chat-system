package main

import (
	"database/sql"
	"log"
	"time"
)

var _db *sql.DB

func initMySQLDB() {
	var err error
	time.Sleep(5 * time.Second) // Wating till the db gets initialized.
	// NOTE: If you want to connect on diff database instead of docker, change the `mysql_db` with the the ip.
	dsn := "root:password@tcp(db:3306)/mysql"
	_db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	// Ensure the database is reachable
	err = _db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
}

// func initRedisDB() *redis.Client {
// 	ctx := context.Background()
// 	addr := "localhost:6379"
// 	client := redis.NewClient(&redis.Options{
// 		Addr: addr,
// 	})
// 	_, err := client.Ping(ctx).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return client
// }
