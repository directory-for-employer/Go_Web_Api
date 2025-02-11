package main

import (
	"fmt"
	"go/web-api/configs"
	"go/web-api/internal/auth"
	"go/web-api/internal/link"
	"go/web-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(database)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is lisen on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
