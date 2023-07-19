package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	handler2 "rest-api-postgres/internal/handler"
	"rest-api-postgres/internal/repository"
	"rest-api-postgres/internal/service"
	"rest-api-postgres/pkg"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error config initialization: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while load env vars: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("Error while creating DB:%v", err)
	}

	repos := repository.New(db)
	services := service.New(repos)
	handler := handler2.New(services)

	srv := new(pkg.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
