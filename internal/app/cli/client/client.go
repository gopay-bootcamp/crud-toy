package client

import (
	"bytes"
	"crud-toy/internal/app/cli/printer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	FindProcs(string) error
	CreateProcs(string, string) error
	DeleteProcs(string) error
	UpdateProcs(string, string, string) error
	ReadAllProcs() error
}

type client struct {
	printer               io.Printer
	connectionTimeoutSecs time.Duration
}

func NewClient() Client {
	return &client{
		printer:               io.PrinterInstance,
		connectionTimeoutSecs: time.Second,
	}
}



type Proc struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (c *client) FindProcs(id string) error {
	client := &http.Client{
		Timeout: c.connectionTimeoutSecs,
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
		Timeout: c.connectionTimeoutSecs,
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
		Timeout: c.connectionTimeoutSecs,
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
		Timeout: c.connectionTimeoutSecs,
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
		Timeout: c.connectionTimeoutSecs,
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
		c.printer.PrintTable(bodyBytes)
	}
	return nil
}
