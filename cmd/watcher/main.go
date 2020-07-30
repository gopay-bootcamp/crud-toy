package main

import (
	"context"
	"fmt"
	"crud-toy/internal/app/service/infra/db/etcd"
)

func main(){
	ctx:= context.Background()
	client := etcd.NewClient()
	defer client.Close()
	watchChan := client.SetWatchOnPrefix(ctx,"key")
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}