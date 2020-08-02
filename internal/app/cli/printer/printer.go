package io

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"encoding/json"
)

type Proc struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type Printer interface {
	Println(string, ...color.Attribute)
	PrintTable([]byte)
}



var PrinterInstance Printer

type printer struct{}

func InitPrinter() {
	if PrinterInstance == nil {
		PrinterInstance = &printer{}
	}
}

func (p *printer) Println(msg string, attr ...color.Attribute) {
	color.New(attr...).Println(msg)
}

func (p *printer) PrintTable(procListBytes []byte) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Author"})
	var procsList []Proc
	json.Unmarshal(procListBytes, &procsList)
	for _, proc := range procsList  {
		t.AppendRows([]table.Row{
			{proc.ID, proc.Name, proc.Author},
		})
		t.AppendSeparator()
	}
	t.Render()
	fmt.Println()
}
