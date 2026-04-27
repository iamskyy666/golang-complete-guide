package main

import (
	"fmt"
	"os"

	"demogithub.com/example/app/models"
	"github.com/olekukonko/tablewriter"
)

// go run ./cmd/api/main.go

func main() {
	var Skyy models.User
	Skyy.Name="Skyy Banerjee"
	fmt.Println("Hello from cmd/api/main.go:",Skyy)
	fmt.Println(Skyy.Name)

	fmt.Println("--------------- tablewriter example ----------------")

	data := [][]string{
		{"Package", "Version", "Status"},
		{"tablewriter", "v0.0.5", "legacy"},
		{"tablewriter", "v1.1.4", "latest"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(data[1:])
	table.Render()
}
	