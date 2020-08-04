package daemon

import (
	"crud-toy/internal/cli/client"
	"errors"
)

var proctorDClient client.Client

func StartClient() {
	proctorDClient = client.NewClient()
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
