package data

type JsonDocument struct {}

func (d JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	println("Send xml data to analytics")
}
