package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	TodoRoute(router, db)

	return router
}
