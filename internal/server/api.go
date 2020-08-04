package server

import (
	"context"
	"crud-toy/internal/server/execution"
	"crud-toy/internal/server/handler"
	"crud-toy/config"
	"crud-toy/internal/server/db/etcd"
	"crud-toy/logger"
	"fmt"
	"net/http"
	"github.com/urfave/negroni"
)

func Start() error {

	appPort := ":" + config.Config().AppPort
	server := negroni.New(negroni.NewRecovery())

	etcdClient := etcd.NewClient()
	exec := execution.NewExec(etcdClient)
	procHandler := handler.NewProcHandler(exec)

	router, err := NewRouter(procHandler)

	if err != nil {
		return err
	}

	defer etcdClient.Close()
	server.UseHandler(router)
	logger.Info("Starting server on port " + string(appPort))
	httpServer := &http.Server{
		Addr:    appPort,
		Handler: server,
	}

	go func() {
		ctx := context.Background()
		client := etcd.NewClient()
		defer client.Close()
		watchChan := client.SetWatchOnPrefix(ctx, "key")
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}
	}()

	if err = httpServer.ListenAndServe(); err != nil {
		logger.Error("HTTP Server Failed ", err)
	}
	logger.Info("Stopped server gracefully")
	return nil
}
