package pkg

const (
	ServerType         = "server"
	PersonComputerType = "personal"
	NotebookType       = "notebook"
)

type IComputer interface {
	GetType() string
	PrintDetail()
}

func New(typeName string) IComputer {
	switch typeName {
	case ServerType:
		return NewServer()
	case NotebookType:
		return NewNotebook()
	case PersonComputerType:
		return NewPersonPC()
	default:
		println("error not found type", typeName)
		return nil
	}
}
