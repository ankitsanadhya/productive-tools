package storage

import (
	"fmt"
	"license/api/model"
	"license/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//InitDB with configuration
func InitDB(configuration config.Configurations) {
	connectionURL := configuration.Database.DBUser + ":" + configuration.Database.DBPassword + "@tcp(" + configuration.Database.HostName + ":" + fmt.Sprintf("%v", configuration.Database.DBPort) + ")/" + "mysql" + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("Connection Url - ", connectionURL)

	var err error
	DB, err = gorm.Open("mysql", connectionURL)
	DB.SingularTable(true)
	DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", configuration.Database.DBName))
	DB.Exec(fmt.Sprintf("USE %s;", configuration.Database.DBName))
	if err != nil {
		fmt.Println("status: ", err)
	}
	// defer DB.Close()

	DB.AutoMigrate(&model.Partner{}, &model.Product{})
	DB.Model(&model.Product{}).AddForeignKey("partner_id", "Partner(partner_id)", "CASCADE", "RESTRICT")
}
