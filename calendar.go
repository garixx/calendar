package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	log.Println(viper.GetString("server.address"))

	log.Fatal("not implemented")
}
