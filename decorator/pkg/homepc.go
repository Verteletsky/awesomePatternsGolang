package pkg

type HomePc struct {
	Cpu     int
	Gpu     int
	Wrapper Wrapper
}

func (h HomePc) GetPrice() float64 {
	return h.Wrapper.GetPrice() + float64(h.Cpu) * float64(h.Gpu)
}
