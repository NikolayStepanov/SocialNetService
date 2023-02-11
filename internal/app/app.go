package app

import (
	httpDelivery "SocialNetHTTPService/internal/delivery/http"
	"SocialNetHTTPService/internal/repository"
	"SocialNetHTTPService/internal/repository/postgres"
	"SocialNetHTTPService/internal/server"
	"SocialNetHTTPService/internal/service"
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	databaseURLKey = "DATABASE_URL"
	portKey        = "PORT"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	wg := sync.WaitGroup{}
	var userDB repository.UsersRep
	dbConfig := os.Getenv(databaseURLKey)
	if dbConfig == "" {
		log.Println("empty env config")
		dbConfig = viper.GetString(databaseURLKey)
	}

	db, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}
	userDB = repository.NewUserDatabase(db)

	userRepository := repository.NewRepositories(userDB)

	services := service.NewServices(service.ServicesDeps{
		Repos: userRepository,
	})

	handlers := httpDelivery.NewHandler(services.Users, services.Friends)
	srv := server.NewServer(handlers.Init())
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = srv.Run(); err != nil {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	log.Print("Server started")
	<-ctx.Done()
	wg.Wait()
	log.Print("Server stopped")
}
