package pkg

import "fmt"

type PersonPC struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonPC() IComputer {
	return &PersonPC{
		Type:   PersonComputerType,
		Core:   8,
		Memory: 16,
		Monitor: true,
	}
}

func (s PersonPC) GetType() string {
	return s.Type
}

func (s PersonPC) PrintDetail() {
	fmt.Printf("Type: %s\nCore: %d\nMemory: %d\nMonitor: %v\n--------\n", s.Type, s.Core, s.Memory, s.Monitor)
}
