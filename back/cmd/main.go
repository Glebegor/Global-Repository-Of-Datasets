package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	handler "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/handler"
	godotenv "github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	handler := new(handler.Handler)
	server := new(grod.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
		}
	}()
	logrus.Printf("%s: Server is running on port %s. BTW", time.Now(), viper.GetString("port"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("%s: Server Shutting Down. Press F... Bro... \n", time.Now())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
