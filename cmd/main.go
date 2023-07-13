package main

import (
	"github.com/spf13/viper"
	"log"
	handler2 "rest-api-postgres/internal/handler"
	"rest-api-postgres/internal/repository"
	"rest-api-postgres/internal/service"
	"rest-api-postgres/pkg"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error config initialization: %v", err)
	}
	repos := repository.New()
	services := service.New(repos)
	handler := handler2.New(services)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
