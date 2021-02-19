package cmd

import (
	"context"
	"errors"
	"go_web_starter/config"
	"go_web_starter/database"
	"go_web_starter/route"
	"go_web_starter/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WebServerStart() {
	config.Init()
	database.Init()
	route.Init()

	server := &http.Server{Addr: ":" + config.Config.AppConfig.Port, Handler: route.Application}

	go func() {
		util.Log.Infof("Application Start Successful...")
		util.Log.Infof("Listen at http://127.0.0.1:" + config.Config.AppConfig.Port)

		var err error

		if config.Config.AppConfig.Https == true {
			err = server.ListenAndServeTLS("./ssl/ssl.pem", "./ssl/ssl.key")
		} else {
			err = server.ListenAndServe()
		}

		if err != nil && errors.Is(err, http.ErrServerClosed) {
			util.Log.Error("Application Start Error:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Log.Error("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		util.Log.Error("Server forced to shutdown:", err)
	}

	util.Log.Info("Server exiting")
}
