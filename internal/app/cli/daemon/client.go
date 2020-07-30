package daemon

import (
	"bytes"
	"crud-toy/internal/app/cli/utility/io"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Proc struct {
	ID     int32  `json:"id"`
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

func (c *client) FindProcs(idString string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	id64, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return err
	}
	id := int32(id64)
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

func (c *client) UpdateProcs(idString string, name string, author string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	id64, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return err
	}
	id := int32(id64)
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

func (c *client) DeleteProcs(idString string) error {
	client := &http.Client{
		Timeout: proctorDClient.connectionTimeoutSecs,
	}
	id64, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return err
	}
	id := int32(id64)
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

func FindProcs(idString string) error {
	return proctorDClient.FindProcs(idString)
}

func CreateProcs(name string, author string) error {
	return proctorDClient.CreateProcs(name, author)
}

func UpdateProcs(idString string, name string, author string) error {
	return proctorDClient.UpdateProcs(idString, name, author)
}

func DeleteProcs(idString string) error {
	return proctorDClient.DeleteProcs(idString)
}
