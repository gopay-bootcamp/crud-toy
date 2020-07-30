package server

import (
	"github.com/gorilla/mux"
	"crud-toy/internal/app/service/infra/execution"
)



func NewRouter(exec execution.Execution) (*mux.Router, error){
	router := mux.NewRouter()
	router.HandleFunc("/create",exec.CreateProc).Methods("POST")
	router.HandleFunc("/read",exec.ReadProcByID).Methods("POST")
	router.HandleFunc("/readAll",exec.ReadAllProc).Methods("GET")
	router.HandleFunc("/update",exec.UpdateProc).Methods("POST")
	router.HandleFunc("/delete",exec.DeleteProc).Methods("POST")
	return router,nil
}






