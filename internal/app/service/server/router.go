package server

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"crud-toy/internal/app/service/infra/db/etcd"
	"github.com/gorilla/mux"
	"math/rand"
	"context"
)


type Proc struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}
func NewRouter() (*mux.Router, error){
	router := mux.NewRouter()

	router.HandleFunc("/create",createProc).Methods("POST")
	router.HandleFunc("/read",readProc).Methods("POST")
	router.HandleFunc("/update",updateProc).Methods("POST")
	router.HandleFunc("/delete",deleteProc).Methods("POST")



	return router,nil
}

var timeout time.Duration = 2 * time.Second


func createProc(w http.ResponseWriter,r *http.Request){
	client := etcd.NewClient()
	defer client.Close()
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	w.Header().Set("Content-Type","application/json")
	var proc Proc
	id := rand.Int31()
	json.NewDecoder(r.Body).Decode(&proc)
	proc.ID=id
	value,err := json.Marshal(proc)
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	fmt.Println(string(value))
	client.PutValue(ctx,string(id),string(value))
	json.NewEncoder(w).Encode(&proc)
	cancel()
}

func readProc(w http.ResponseWriter,r *http.Request){
	client := etcd.NewClient()
	defer client.Close()
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	w.Header().Set("Content-Type","application/json")
	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	id := proc.ID
	gr,err :=client.GetValue(ctx,string(id))
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(string(gr.Kvs[0].Value))
	cancel()
}

func updateProc(w http.ResponseWriter,r *http.Request){
	client := etcd.NewClient()
	defer client.Close()
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	w.Header().Set("Content-Type","application/json")
	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	value,err := json.Marshal(proc)
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	fmt.Println(string(value))
	client.PutValue(ctx,string(proc.ID),string(value))
	json.NewEncoder(w).Encode(&proc)
	cancel()
}

func deleteProc(w http.ResponseWriter,r *http.Request){
	client := etcd.NewClient()
	defer client.Close()
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	w.Header().Set("Content-Type","application/json")
	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	id := proc.ID
	err:=client.DeleteKey(ctx,string(id))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode("value deleted")
	cancel()
}

