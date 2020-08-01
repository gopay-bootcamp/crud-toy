package execution

import (
	"strconv"
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"crud-toy/internal/app/service/infra/db/etcd"
	crypto_rand "crypto/rand"
	math_rand "math/rand"
	"context"
	"encoding/binary"
)

var (
	timeout time.Duration = 2 * time.Second
	prefix string = "key"

)

type Execution interface {
	CreateProc(w http.ResponseWriter,r *http.Request)
	ReadProcByID(w http.ResponseWriter,r *http.Request)
	ReadAllProc(w http.ResponseWriter,r *http.Request)
	UpdateProc(w http.ResponseWriter,r *http.Request)
	DeleteProc(w http.ResponseWriter,r *http.Request)
	CloseDbClient()
}

type execution struct {
	client etcd.EtcdClient
	ctx context.Context
	cancel context.CancelFunc
}

type Proc struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}

func GetDbClient() Execution{
	client := etcd.NewClient()
	return &execution{
		client:client,
	}
}

func (e *execution) CreateProc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)
	var proc Proc
	var b [8]byte
	crypto_rand.Read(b[:])
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	json.NewDecoder(r.Body).Decode(&proc)
	proc.ID=strconv.FormatInt(math_rand.Int63(),10)
	value,err := json.Marshal(proc)
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	fmt.Println(string(value))
	e.client.PutValue(ctx,fmt.Sprintf("key_%s",proc.ID),string(value))
	json.NewEncoder(w).Encode(&proc)
	cancel()
}

func (e *execution) ReadProcByID(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)

	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	id := proc.ID
	gr,err :=e.client.GetValue(ctx,fmt.Sprintf("key_%s",id))
	if err !=nil{
		w.Write([]byte(err.Error()))
	}

	if len(gr.Kvs)==0 {
	json.NewEncoder(w).Encode(string("no value present"))
	} else {
	w.Write(gr.Kvs[0].Value)
	}
	cancel()

}

func (e *execution) ReadAllProc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)

	gr,err :=e.client.GetAllValueWithPrefix(ctx,prefix)
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	var procs []Proc
	for _,kv := range gr.Kvs{
		proc := Proc{}
		str := string(kv.Value)
		json.Unmarshal([]byte(str),&proc)
		procs = append(procs,proc)
	}
	//value,err := json.Marshal(procs)

	json.NewEncoder(w).Encode(procs)
	cancel()

}

func (e *execution) UpdateProc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)

	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	value,err := json.Marshal(proc)
	if err !=nil{
		w.Write([]byte(err.Error()))
	}
	e.client.PutValue(ctx,fmt.Sprintf("key_%s",proc.ID),string(value))
	json.NewEncoder(w).Encode(&proc)
	cancel()

}

func (e *execution) DeleteProc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ctx,cancel := context.WithTimeout(context.Background(),timeout)

	var proc Proc
	json.NewDecoder(r.Body).Decode(&proc)
	id := proc.ID
	err:=e.client.DeleteKey(ctx,fmt.Sprintf("key_%s",id))
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode("value deleted")
	cancel()
}

func (e *execution) CloseDbClient(){
	defer e.client.Close()
}