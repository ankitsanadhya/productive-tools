package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfiguration() (configuration Configurations, m Messages) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("./config")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	viper.SetConfigName("sql")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.MergeInConfig()

	viper.SetConfigName("messages-en")
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	viper.MergeInConfig()

	err = viper.Unmarshal(&m)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Database is\t", m.Data)
	fmt.Printf(fmt.Sprintf(m.Data, 1, 66))
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Create Database \t\t", viper.GetString("create.database"))
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))

	return
}

func InitConfig() {
	var config Configurations
	appConfig, _ := loadConfigurations("./config/config.yml")
	config = appConfig.(Configurations)
	fmt.Println(config)

	//sqlConfig, _ := loadConfigurations("./config/sql.yml")

}

func loadConfigurations(filePath string) (intfase interface{}, err error) {
	// configsPath := constant.ProductPath() + constant.LocalSeperator()
	// os.MkdirAll(configsPath, os.ModePerm)
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&intfase)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	return
}
