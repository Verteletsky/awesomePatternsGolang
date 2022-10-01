package pkg

const (
	AsusCollectorType = "asus"
	HPCollectorType   = "hp"
)

type ICollector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGPU()
	GetComputer() Computer
}

func GetCollector(collectorType string) ICollector {
	switch collectorType {
	case AsusCollectorType:
		return &AsusCollector{}
	case HPCollectorType:
		return &HPCollector{}
	default:
		return nil
	}
}
