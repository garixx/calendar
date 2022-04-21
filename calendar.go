package main

import (
	"context"
	"github.com/garixx/calendar/frontend"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func main() {
	readConfig()
	f, err := frontend.NewFrontEnd(viper.GetString("frontend"))
	if err != nil {
		log.Fatal(err)
	}
	f.Start(context.TODO())
	//r := mux.NewRouter()
	//r.HandleFunc("/register", handlers.RegistrationHandler).Methods("POST")

	//srv := &http.Server{
	//	Addr:         "0.0.0.0:" + viper.GetString("server.address"),
	//	WriteTimeout: time.Second * 15,
	//	ReadTimeout:  time.Second * 15,
	//	IdleTimeout:  time.Second * 60,
	//	Handler:      r,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Println(err)
	//	}
	//}()

	//c := make(chan os.Signal, 1)
	//// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	//// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	//signal.Notify(c, os.Interrupt)
	//
	//// Block until we receive our signal.
	//<-c
	//gracefulShutdown(srv)
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
