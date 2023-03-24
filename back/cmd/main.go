package main

import (
	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializating configs: %s", err.Error())
	}
	server := new(grod.Server)
	if err := server.Run(viper.GetString("port")); err != nil {
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
