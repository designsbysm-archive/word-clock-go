// TODO: add documentation
package main

import (
	"github.com/designsbysm/word-clock-server-go/api"
)

func main() {
	if err := config(); err != nil {
		panic(err)
	}

	if err := loggers(); err != nil {
		panic(err)
	}

	// run each server
	api.Serve()
}
