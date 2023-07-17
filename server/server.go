package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/server/router"
	utils "github.com/delta/arcadia-backend/utils"
)

func Run() {
	config := config.GetConfig()

	// Initialize all the routes
	router.Init()

	utils.Logger.Println("Server started")

	server := http.Server{
		Addr:    config.Host + ":" + strconv.FormatUint(uint64(config.Port), 10),
		Handler: router.Router,
	}

	// To Gracefully shutdown https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	utils.Logger.Error("Shutdown Server ...")

	// Timeout of 2s
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.Logger.Errorf("Server Shutdown:%s", err)
	}

	<-ctx.Done()
	utils.Logger.Errorf("Timeout of %ds\n", 2)

	utils.Logger.Error("Server exiting")
}
