package cmd

import (
	"context"
	"errors"
	"go_web_starter/config"
	"go_web_starter/dao"
	"go_web_starter/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WebServerStart() {
	config.Init()
	dao.Init()
	route.Init()

	server := &http.Server{Addr: ":8080", Handler: route.Application}

	go func() {
		log.Print("Application Start Successful...")
		log.Print("Listen at http://127.0.0.1:8080")

		err := server.ListenAndServe()

		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Print("Application Start Error:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
