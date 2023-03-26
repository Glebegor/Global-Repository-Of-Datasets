package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	handler "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/handler"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/repository"
	"github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/service"
	godotenv "github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.ConnectToDB(repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initializate database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(grod.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
		}
	}()
	logrus.Printf("%s: Server is running on port %s. BTW", time.Now(), viper.GetString("port"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("%s: Server Shutting Down. Press F... Bro... \n", time.Now())

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Printf("Error while server was shutdowning: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
