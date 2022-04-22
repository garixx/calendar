package frontend

import (
	"context"
	"github.com/garixx/calendar/internals/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type restFrontEnd struct{}

func (f restFrontEnd) Start(ctx context.Context) error {
	log.Println("it's me, rest frontend")
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegistrationHandler).Methods("POST")
	srv := &http.Server{
		Addr:         viper.GetString("server.address") + ":" + viper.GetString("server.port"),
		WriteTimeout: time.Second * viper.GetDuration("server.timeout.write"),
		ReadTimeout:  time.Second * viper.GetDuration("server.timeout.read"),
		IdleTimeout:  time.Second * viper.GetDuration("server.timeout.idle"),
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	gracefulShutdown(srv)
	return nil
}

func gracefulShutdown(srv *http.Server) {
	//TODO: add DB closing and other necessary actions
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
