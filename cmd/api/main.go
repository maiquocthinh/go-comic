package main

import (
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/internal/server"
	"github.com/maiquocthinh/go-comic/pkg/db/mysql"
	"github.com/maiquocthinh/go-comic/pkg/db/redis"
	"log"
)

func main() {
	log.Println("Starting API server")

	// load config
	cfg, err := config.NewConfig("./config/config-local")
	if err != nil {
		log.Fatalln("Load config fail!")
	}

	// new mysql
	mysqlDB, err := mysql.NewMysqlDB(&cfg.MySQL)
	if err != nil {
		log.Fatalf("MySQL init: %s", err)
	} else {
		log.Printf("MySQL connected, Status: %#v\n", mysqlDB.Stats())
	}
	defer mysqlDB.Close()

	// new redis client
	redisClient := redis.NewRedisClient(&cfg.Redis)

	// start server
	s := server.NewServer(cfg, mysqlDB, redisClient)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
