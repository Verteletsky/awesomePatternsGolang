package pkg

type Factory struct {
	Collector ICollector
}

func NewFactory(collector ICollector) *Factory {
	return &Factory{Collector: collector}
}

func (f *Factory) SetCollector(collector ICollector) {
	f.Collector = collector
}

func (f *Factory) CreateComputer() Computer {
	f.Collector.SetCore()
	f.Collector.SetMemory()
	f.Collector.SetBrand()
	f.Collector.SetGPU()
	f.Collector.SetMonitor()
	return f.Collector.GetComputer()
}
