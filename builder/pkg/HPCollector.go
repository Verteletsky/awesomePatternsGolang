package pkg

type HPCollector struct {
	Core    int
	Brand   string
	Memory  int
	Monitor int
	GPU     int
}

func (collector *HPCollector) SetCore() {
	collector.Core = 4
}

func (collector *HPCollector) SetBrand() {
	collector.Brand = "HP"
}

func (collector *HPCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HPCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HPCollector) SetGPU() {
	collector.GPU = 1
}

func (collector *HPCollector) GetComputer() Computer {
	 return Computer{
		 Core:    collector.Core,
		 Brand:   collector.Brand,
		 Memory:  collector.Memory,
		 Monitor: collector.Monitor,
		 GPU:     collector.GPU,
	 }
}

