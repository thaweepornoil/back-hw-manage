package connections

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Error error

func Conection_hw() (*gorm.DB, error) {

	dsn := "root:root@tcp(127.0.0.1:8889)/hw-manage?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db, err

}

func Conection() {

	dsn := "root:root@tcp(127.0.0.1:8889)/hw-manage?charset=utf8mb4&parseTime=True&loc=Local"
	DB, Error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if Error != nil {
		panic("failed to connect database")
	}
	fmt.Println(DB)
}
