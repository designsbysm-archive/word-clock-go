package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	// create
	router := gin.New()
	AddRoute(router)

	address := viper.GetString("api.address")

	// notify
	if viper.GetBool("gin.release") {
		timber.Info(fmt.Sprintf("API: listening on %s (%s)", address, "HTTP"))

		if err := router.Run(address); err != nil {
			timber.Error("API:", err)
		}
	}

	// wait for ^c
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// close
	fmt.Println("")
	timber.Info("API: closed")
}
