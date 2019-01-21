package main

import (
	"github.com/spf13/viper"
)

type serverConfig struct {
	Address   string
	MysqlHost string
	MysqlPort string
	MysqlUser string
	MysqlPass string
	MysqlDb   string
}

var (
	configuration *serverConfig
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	configuration = &serverConfig{
		Address:   viper.GetString("server.address"),
		MysqlHost: viper.GetString("mysql.host"),
		MysqlPort: viper.GetString("mysql.port"),
		MysqlUser: viper.GetString("mysql.user"),
		MysqlPass: viper.GetString("mysql.pass"),
		MysqlDb:   viper.GetString("mysql.db"),
	}
}
