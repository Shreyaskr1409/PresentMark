package routes

import (
	"log"

	"github.com/Shreyaskr1409/PresentMark/handlers"
	"github.com/gorilla/mux"
)

func HandleFileRoutes(l *log.Logger, parentRouter *mux.Router) {
	fileHandler := handlers.InitFileHandler(l)

	fileRoute := parentRouter.PathPrefix("/api/v1/files").Subrouter()
	fileRoute.HandleFunc("/{id}", fileHandler.GetFile).Methods("GET")
	fileRoute.HandleFunc("/{id}", fileHandler.CreateFile).Methods("POST")
	fileRoute.HandleFunc("/{id}", fileHandler.DeleteFile).Methods("DELETE")
	fileRoute.HandleFunc("/{id}", fileHandler.UpdateFile).Methods("PATCH")
}
