package src

import (
	"encoding/xml"
	"errors"
)

type XmlReader struct{}

func (XmlReader) Read(xmlFile []byte) (Api, error) {
	// we initialize our Api variable
	var api Api

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	unmarshalErr := xml.Unmarshal(xmlFile, &api)

	if unmarshalErr != nil {
		return Api{}, errors.New(unmarshalErr.Error())
	}

	return api, nil
}

type Api struct {
	XMLName   xml.Name   `xml:"api"`
	Functions []Function `xml:"function"`
}

type Function struct {
	XMLName xml.Name `xml:"function"`
	Name    string   `xml:"name,attr"`
	Inputs  []Input  `xml:"input"`
	Outputs []Output `xml:"output"`
}

type Input struct {
	XMLName     xml.Name `xml:"input"`
	Name        string   `xml:"name,attr"`
	Type        string   `xml:"type,attr"`
	Optional    bool     `xml:"optional,attr"`
	Description string   `xml:"desc,attr"`
}

type Output struct {
	XMLName     xml.Name `xml:"output"`
	Name        string   `xml:"name,attr"`
	Type        string   `xml:"type,attr"`
	Description string   `xml:"desc,attr"`
}
