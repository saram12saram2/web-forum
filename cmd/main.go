package main

import (
	"log"
	"web-forum/configs"
	"web-forum/internal/handler"
	"web-forum/internal/server"
	"web-forum/internal/service"
	"web-forum/internal/store"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.NewSqlite3(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := store.New(db)
	service := service.New(store)
	handler := handler.New(service)
	srv := new(server.Server)

	if err := srv.Run(config.Port, handler.InitRouters()); err != nil {
		log.Fatalf("Error in main: %s", err)
	}
}
