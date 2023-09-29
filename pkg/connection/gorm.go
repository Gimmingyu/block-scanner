package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection(user, pass, database, host string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, pass, host, database)
	client, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return client
}
