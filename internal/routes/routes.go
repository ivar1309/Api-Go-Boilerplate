package routes

import (
	"github.com/ivar1309/Api-Go-Boilerplate/internal/controllers"
	"github.com/ivar1309/Api-Go-Boilerplate/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/public", controllers.PublicEndpoint).Methods("GET")
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")

	protected := r.PathPrefix("/api/protected").Subrouter()
	protected.Use(middleware.JWTAuthMiddleware)
	protected.HandleFunc("", controllers.ProtectedEndpoint).Methods("GET")

	return r
}
