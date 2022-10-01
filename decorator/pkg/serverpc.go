package pkg

type ServerPC struct {
	Cpu     int
	Memory  int
	Wrapper Wrapper
}

func (s ServerPC) GetPrice() float64 {
	return s.Wrapper.GetPrice() * float64(s.Cpu) * float64(s.Memory)
}
