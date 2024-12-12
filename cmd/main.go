package main

import (
	"log"
	"net/http"

	"github.com/Sukanta2002/todo-api-golang/db"
	"github.com/Sukanta2002/todo-api-golang/middleware"
	"github.com/Sukanta2002/todo-api-golang/routes"
)

func main() {
	dbConn := db.InitDb()
	defer dbConn.Close()

	route := routes.SetupRoutes(dbConn)
	route.Use(middleware.LoggingMiddleware)

	log.Println("Server starting at port 8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}
