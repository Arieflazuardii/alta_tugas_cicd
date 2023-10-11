package main

import (
	"praktikum/config"
	m "praktikum/middleware"
	"praktikum/routes"
)


func main() {
  config.LoadEnv()
  config.InitDB()
  // create a new echo instance
  e := routes.New()
  m.LogMiddleware(e)
  // Route / to handler function
  e.Logger.Fatal(e.Start(":8080"))
}