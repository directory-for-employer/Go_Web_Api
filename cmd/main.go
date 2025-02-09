package main

import (
	"fmt"
	"go/web-api/configs"
	"go/web-api/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
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
