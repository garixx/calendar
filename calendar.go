package main

import (
	"context"
	"github.com/garixx/calendar/internals/frontend"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func main() {
	readConfig()
	f, err := frontend.New(viper.GetString("frontend"))
	if err != nil {
		log.Fatal(err)
	}
	f.Start(context.TODO())
}

func readConfig() {
	// make qa environment default
	_ = viper.BindEnv("env")
	viper.SetDefault("env", "QA")

	viper.SetConfigFile(`./configs/` + strings.ToLower(viper.GetString("env")) + `/config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//func gracefulShutdown(srv *http.Server) {
//	//TODO: add DB closing and other necessary actions
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
//	defer cancel()
//	srv.Shutdown(ctx)
//	log.Println("shutting down")
//	os.Exit(0)
//}
