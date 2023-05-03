package connection

import (
	"gin_unsplash/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(conf config.MySQL) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(conf.DSN()))
	if err != nil {
		return nil, err
	}
	return db, nil
}
