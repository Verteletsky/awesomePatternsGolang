package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() IComputer {
	return &Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetail() {
	fmt.Printf("Type: %s\nCore: %d\nMemory: %d\n--------\n", s.Type, s.Core, s.Memory)
}
