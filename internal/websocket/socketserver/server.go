package socketserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/coreos/etcd/clientv3"
)

var race sync.RWMutex

func Start(watchChan clientv3.WatchChan) {
	listener := SetupListener("tcp", "localhost:49152")
	writeChans := make(map[string]chan string)
	go func() {
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				race.RLock()
				for _, wChan := range writeChans {
					wChan <- fmt.Sprintf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
				}
				race.RUnlock()
			}
		}
	}()
	for {
		connectionWithClient := SetupConnection(listener)
		defer connectionWithClient.Close()
		go SetupReaderAndWriter(connectionWithClient, writeChans)
	}
}

func SetupListener(network, address string) net.Listener {
	listener, listenerErr := net.Listen(network, address)
	if listenerErr != nil {
		log.Fatal(listenerErr)
	}
	log.Printf("Server is listening at network address: %v\n", listener.Addr())
	return listener
}

func SetupConnection(clientListener net.Listener) net.Conn {
	log.Println("Waiting for client to dial...")
	connectionWithClient, connectionWithClientErr := clientListener.Accept()
	if connectionWithClientErr != nil {
		log.Fatal(connectionWithClientErr)
	}
	log.Printf("Establishing connection with client at network address: %v", connectionWithClient.RemoteAddr())
	return connectionWithClient
}
func acceptMessageFromClient(connectionWithClient net.Conn, writeChan chan string, writeChans map[string](chan string)) {
	for {
		reader := bufio.NewReader(connectionWithClient)
		dataFromClient, dataFromClientError := reader.ReadString('\n')
		if dataFromClientError.Error() == "EOF" {
			removeClientChannel(connectionWithClient, writeChan, writeChans)
			return
		}
		if dataFromClientError != nil {
			log.Fatal(dataFromClientError)
			return
		}
		if strings.TrimSpace(string(dataFromClient)) == "STOP" {
			removeClientChannel(connectionWithClient, writeChan, writeChans)
			return
		}
		fmt.Print("From client -> ", string(dataFromClient))
	}
}
func setWatchForClient(connectionWithClient net.Conn, writeChan chan string) {
	for message := range writeChan {
		if strings.TrimSpace(message) == "STOP" {
			return
		}
	}
}
func SetupReaderAndWriter(connectionWithClient net.Conn, writeChans map[string]chan string) {
	newChan := createClientChannel(connectionWithClient, writeChans)
	log.Println("Number of channels: ", len(writeChans))
	log.Println("Setting watcher for client at: ", connectionWithClient.RemoteAddr())
	go acceptMessageFromClient(connectionWithClient, newChan, writeChans)
	setWatchForClient(connectionWithClient, newChan)
	log.Print("Closing connection with client at: ", connectionWithClient.RemoteAddr())
}

func removeClientChannel(connectionWithClient net.Conn, writeChan chan string, writeChans map[string]chan string) {
	fmt.Print("Deleting the channel")
	close(writeChan)
	race.Lock()
	delete(writeChans, connectionWithClient.RemoteAddr().String())
	race.Unlock()
}

func createClientChannel(connectionWithClient net.Conn, writeChans map[string]chan string) chan string {
	newChan := make(chan string)
	race.Lock()
	writeChans[connectionWithClient.RemoteAddr().String()] = newChan
	race.Unlock()
	return newChan
}
