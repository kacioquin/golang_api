package main

import (
	"fmt"

	"github.com/kacioquin/golang_api/handlers"
	"github.com/kacioquin/golang_api/middlewares"
	"github.com/kacioquin/golang_api/router"
)

func main() {
	fmt.Println("Func main")
	middlewares.Middleware()
	router.Router()
	handlers.Handler()
}
