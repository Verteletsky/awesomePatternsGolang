package pkg

import "fmt"

type Computer struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	GPU     int
}

func (c Computer) Print() {
	fmt.Printf("Brand: %s, Core: %d, Memory: %d, GPU: %d, Monitor: %d", c.Brand, c.Core, c.Memory, c.GPU, c.Monitor)
}
