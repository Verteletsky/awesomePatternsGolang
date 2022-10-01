package data_service

type IAnalyticsDataService interface {
	SendXmlData()
}

type XmlDocument struct {}

func (x XmlDocument) SendXmlData() {
	println("Send xml document")
}

