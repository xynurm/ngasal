package routes

import (
	"ngasal/handlers"
	"ngasal/pkg/middleware"
	"ngasal/pkg/mysql"
	"ngasal/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/category", middleware.Auth(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/categories", middleware.Auth(h.FindCategories)).Methods("GET")
}
