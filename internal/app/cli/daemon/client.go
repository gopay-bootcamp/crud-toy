package daemon

import (
	"bytes"
	"crud-toy/internal/app/cli/utility/io"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Proc struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type Client interface {
	FindProcs(string) error
	CreateProcs(string) error
	DeleteProcs(string) error
	UpdateProcs(string) error
}

type client struct {
	printer               io.Printer
	connectionTimeoutSecs time.Duration
}

var proctorDClient *client

func StartClient(printer io.Printer) {
	proctorDClient = &client{
		printer:               printer,
		connectionTimeoutSecs: time.Second,
	}
}

func (c *client) FindProcs(id string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	requestBody, err := json.Marshal(&Proc{
		ID: id,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:3000/read", bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
	return nil
}

func (c *client) CreateProcs(name string, author string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	requestBody, err := json.Marshal(&Proc{
		Name:   name,
		Author: author,
	})
	req, err := http.NewRequest("POST", "http://localhost:3000/create", bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
	return nil
}

func (c *client) UpdateProcs(id string, name string, author string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	requestBody, err := json.Marshal(&Proc{
		ID:     id,
		Name:   name,
		Author: author,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:3000/update", bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
	return nil
}

func (c *client) DeleteProcs(id string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	requestBody, err := json.Marshal(&Proc{
		ID: id,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:3000/delete", bytes.NewReader(requestBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
	return nil
}

func (c *client) ReadAllProcs() error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	req, err := http.NewRequest("GET", "http://localhost:3000/readAll", nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
	return nil
}

func FindProcs(args []string) error {
	if len(args) < 1 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	return proctorDClient.FindProcs(id)
}

func CreateProcs(args []string) error {
	if len(args) < 2 {
		return errors.New("Invalid argument error")
	}
	name := args[0]
	author := args[1]
	return proctorDClient.CreateProcs(name, author)
}

func UpdateProcs(args []string) error {
	if len(args) < 3 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	name := args[1]
	author := args[2]
	return proctorDClient.UpdateProcs(id, name, author)
}

func DeleteProcs(args []string) error {
	if len(args) < 1 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	return proctorDClient.DeleteProcs(id)
}

func ReadAllProcs() error {
	return proctorDClient.ReadAllProcs()
}
