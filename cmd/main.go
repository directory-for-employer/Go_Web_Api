package main

import (
	"fmt"
	"go/web-api/configs"
	"go/web-api/internal/auth"
	"go/web-api/internal/link"
	"go/web-api/internal/user"
	"go/web-api/pkg/db"
	"go/web-api/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)

	// Service
	authService := auth.NewAuthService(userRepository)

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	user.NewUserHandler(router, user.UserHandlerDeps{
		UserRepository: userRepository,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("server is lisen on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
