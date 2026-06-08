package router

import (
	"database/sql"
	"mastering-actions/internal/config"
	"mastering-actions/internal/handlers"
	"mastering-actions/internal/middleware"
	"mastering-actions/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB, cfg *config.Config) *mux.Router {
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()

	// Apply global middleware (now compatible with gorilla/mux)
	r.Use(middleware.CORS)
	r.Use(middleware.SecurityHeaders)
	r.Use(middleware.Logging)
	r.Use(middleware.RateLimit)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// User routes - no need to wrap with middleware here since global middleware applies
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}
