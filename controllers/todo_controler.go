package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sukanta2002/todo-api-golang/models"
	"github.com/Sukanta2002/todo-api-golang/utils"
	"github.com/gorilla/mux"
)

type TodoController struct {
	DB *sql.DB
}

func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	data := r.Body

	if err := json.NewDecoder(data).Decode(&todo); err != nil {
		fmt.Println(err)
		utils.ApiError(w, http.StatusBadRequest, "Invalid json")
		return
	}
	if todo.Title == "" && todo.Description == "" {
		utils.ApiError(w, http.StatusBadRequest, "Enter Titel and discription")
		return
	}

	_, err := c.DB.Exec("INSERT INTO todos (title, description) VALUES (?, ?)", todo.Title, todo.Description)
	if err != nil {
		fmt.Println(err)
		utils.ApiError(w, http.StatusBadRequest, "SQL error")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]string{"title": todo.Title, "description": todo.Description})

}

func (c *TodoController) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Println(err)
		utils.ApiError(w, http.StatusBadRequest, "SQL error")
		return
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
		todos = append(todos, todo)
	}

	utils.RespondJSON(w, http.StatusOK, todos)

}

func (c *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	if id == "" {
		utils.ApiError(w, http.StatusBadRequest, "ID required")
		return
	}
	var todo models.Todo
	err := c.DB.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)

	if err == sql.ErrNoRows {
		utils.ApiError(w, http.StatusNotFound, "NO TODO of this id")
		return
	} else if err != nil {
		utils.ApiError(w, http.StatusBadRequest, "SQL error")
		return
	}

	utils.RespondJSON(w, http.StatusOK, todo)
}

// DeleteTodo deletes a TODO item by its ID
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		utils.ApiError(w, http.StatusBadRequest, "missing ID parameter")
		return
	}

	// Convert the ID to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ApiError(w, http.StatusBadRequest, "invalid ID format")
		return
	}

	// Execute the DELETE query
	result, err := c.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		utils.ApiError(w, http.StatusInternalServerError, "failed to delete todo")
		return
	}

	// Check if any row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.ApiError(w, http.StatusInternalServerError, "could not verify deletion")
		return
	}
	if rowsAffected == 0 {
		utils.ApiError(w, http.StatusNotFound, "todo not found")
		return
	}

	// Respond with success
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
}
