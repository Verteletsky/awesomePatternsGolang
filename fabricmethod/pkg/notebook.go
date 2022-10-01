package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() IComputer {
	return &Notebook{
		Type:   NotebookType,
		Core:   8,
		Memory: 32,
		Monitor: true,
	}
}

func (s Notebook) GetType() string {
	return s.Type
}

func (s Notebook) PrintDetail() {
	fmt.Printf("Type: %s\nCore: %d\nMemory: %d\nMonitor: %v\n--------\n", s.Type, s.Core, s.Memory, s.Monitor)
}
