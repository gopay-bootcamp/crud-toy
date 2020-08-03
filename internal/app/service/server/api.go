package server

import (
	"context"
	"crud-toy/internal/app/service/execution"
	"crud-toy/internal/app/service/handler"

	"crud-toy/internal/app/service/infra/config"
	"crud-toy/internal/app/service/infra/db/etcd"
	"crud-toy/internal/app/service/infra/logger"
	"fmt"
	"net/http"

	// "os"
	// "os/signal"
	// "syscall"
	"github.com/urfave/negroni"
)

func Start() error {

	appPort := ":" + config.Config().AppPort
	server := negroni.New(negroni.NewRecovery())
	etcdClient := etcd.NewClient()
	exec := execution.NewExec(etcdClient)
	procHanlder := handler.NewProcHandler(exec)

	router, err := NewRouter(procHanlder)
	defer exec.CloseDbClient()
	if err != nil {
		return err
	}
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

	// go func() {
	// 	sigint := make(chan os.Signal, 1)

	// 	signal.Notify(sigint, os.Interrupt)
	// 	signal.Notify(sigint, syscall.SIGTERM)

	// 	<-sigint

	// 	if shutdownErr := httpServer.Shutdown(context.Background()); shutdownErr != nil {
	// 		logger.Error("Received an Interrupt Signal", shutdownErr)
	// 	}
	// }()

	if err = httpServer.ListenAndServe(); err != nil {
		logger.Error("HTTP Server Failed ", err)
	}
	logger.Info("Stopped server gracefully")
	return nil
}
