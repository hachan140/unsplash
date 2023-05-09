package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type MySQL struct {
	Host     string `envconfig:"MYSQL_HOST"`
	Port     int    `envconfig:"MYSQL_PORT"`
	Username string `envconfig:"MYSQL_USERNAME"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	Database string `envconfig:"MYSQL_DATABASE"`
}

func (m MySQL) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
}

func NewMySQLConfig() MySQL {
	var mySQL MySQL
	if err := envconfig.Process("", &mySQL); err != nil {
		log.Fatal("httperror when parsing MySQLConfig, httperror: ", err)
	}
	return mySQL
}
