package main

import "awesomeProject/fabricmethod/pkg"

var types = []string{
	pkg.PersonComputerType,
	pkg.NotebookType,
	pkg.ServerType,
	"mobile",
}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetail()
	}
}
