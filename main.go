package main

import (
	"github.com/teddy-137/task_manager_api/handlers"
)

func main() {
	handlers.StartDB()
	handlers.Start()
}
