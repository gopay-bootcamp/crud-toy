package server

import (
	"crud-toy/internal/app/service/handler"

	"github.com/gorilla/mux"
)

func NewRouter(procHandler handler.Handler) (*mux.Router, error) {
	router := mux.NewRouter()

	router.HandleFunc("/create", procHandler.CreateProc).Methods("POST")
	router.HandleFunc("/read", procHandler.ReadProcByID).Methods("POST")
	router.HandleFunc("/readAll", procHandler.ReadAllProc).Methods("GET")
	router.HandleFunc("/update", procHandler.UpdateProc).Methods("POST")
	router.HandleFunc("/delete", procHandler.DeleteProc).Methods("POST")
	return router, nil
}
