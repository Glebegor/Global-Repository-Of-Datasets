package main

import (
	"time"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	handler "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/handler"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializating configs: %s", err.Error())
	}
	handler := new(handler.Handler)
	server := new(grod.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
