package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Sukanta2002/todo-api-golang/controllers"
	"github.com/gorilla/mux"
)

func TodoRoute(router *mux.Router, db *sql.DB) {

	todoControler := &controllers.TodoController{DB: db}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hiiiiiiii")
	})

	router.HandleFunc("/todo", todoControler.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", todoControler.GetAllTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", todoControler.GetTodoByID).Methods("GET")
	router.HandleFunc("/todo/{id}", todoControler.DeleteTodo).Methods("DELETE")

}
