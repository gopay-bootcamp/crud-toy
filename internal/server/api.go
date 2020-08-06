package server

import (
	"context"
	"crud-toy/config"
	"crud-toy/internal/server/db/etcd"
	"crud-toy/internal/server/execution"
	"crud-toy/internal/server/handler"
	"crud-toy/internal/websocket/socketserver"
	"crud-toy/logger"
	"net/http"

	"github.com/urfave/negroni"
)

func Start() error {

	ctx := context.Background()
	appPort := ":" + config.Config().AppPort
	server := negroni.New(negroni.NewRecovery())

	etcdClient := etcd.NewClient()
	defer etcdClient.Close()

	watchChan := etcdClient.SetWatchOnPrefix(ctx, "key")
	go socketserver.Start(watchChan)

	exec := execution.NewExec(etcdClient)
	procHandler := handler.NewProcHandler(exec)

	router, err := NewRouter(procHandler)

	if err != nil {
		return err
	}

	server.UseHandler(router)
	logger.Info("Starting server on port " + string(appPort))
	httpServer := &http.Server{
		Addr:    appPort,
		Handler: server,
	}

	if err = httpServer.ListenAndServe(); err != nil {
		logger.Error("HTTP Server Failed ", err)
	}
	logger.Info("Stopped server gracefully")
	return nil
}
