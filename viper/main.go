package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	//NAME OF THE FOLDER.IN HERE THAT IS config
	viper.AddConfigPath("config")
	//NAME OF FOLE,IN HERE THAT IS config
	viper.SetConfigName("config")
	// READ THE CONFIG
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error is :", err)
	}
	//GET THE VALUE
	port := viper.Get("server.port")
	name := viper.Get("server.name")
	data := viper.Get("server.database")
	fmt.Println(port, ":", name, ":", data)
}
