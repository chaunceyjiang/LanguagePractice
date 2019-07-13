package protocol

import (
	"encoding/xml"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

type (
	Attr struct {
		Text   string `xml:",chardata"`
		Op     string `xml:"op,attr"`
		Value  string `xml:"value,attr"`
		Type   string `xml:"type,attr"`
		Item1  string `xml:"item1,attr"`
		Item2  string `xml:"item2,attr"`
		Start  string `xml:"start,attr"`
		File   string `xml:"file,attr"`
		Result string `xml:"result,attr"`
		End    string `xml:"end,attr"`
	}
	Field struct {
		Attr `xml:""`
	}

	Basic struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name"`
		DisplayName string `xml:"display_name"`
		Decodes     struct {
			Text   string   `xml:",chardata"`
			Decode []string `xml:"decode"`
		} `xml:"decodes"`
	}
	If struct {
		Attr `xml:""`
		//Text  string  `xml:",chardata"`
		//Op    string  `xml:"op,attr"`
		//Item1 string  `xml:"item1,attr"`
		//Value string  `xml:"value,attr"`
		//Item2 string  `xml:"item2,attr"`
		Fields []Field `xml:"field"`
	}
	Else struct {
		Text   string  `xml:",chardata"`
		Fields []Field `xml:"field"`
	}
	Normalize struct {
		Text   string  `xml:",chardata"`
		Name   string  `xml:"name,attr"`
		If     []If    `xml:"if"`
		Else   Else    `xml:"else"`
		Fields []Field `xml:"field"`
	}
	Protocol struct {
		XMLName xml.Name `xml:"protocol"`
		Text    string   `xml:",chardata"`
		Basic   Basic    `xml:"basic"`
		Prepare struct {
			Text   string  `xml:",chardata"`
			Fields []Field `xml:"field"`
		} `xml:"prepare"`
		Normalizes struct {
			Text      string      `xml:",chardata"`
			Normalize []Normalize `xml:"normalize"`
		} `xml:"normalizes"`
		Traces struct {
			Text  string `xml:",chardata"`
			Trace []struct {
				Text   string  `xml:",chardata"`
				Name   string  `xml:"name,attr"`
				Fields []Field `xml:"field"`
			} `xml:"trace"`
		} `xml:"traces"`
		TraceDisp struct {
			Text   string  `xml:",chardata"`
			Fields []Field `xml:"field"`
		} `xml:"trace_disp"`
		Dimensions struct {
			Text      string `xml:",chardata"`
			Dimension []struct {
				Text        string `xml:",chardata"`
				Name        string `xml:"name,attr"`
				DisplayName string `xml:"display_name,attr"`
				CsvFile     string `xml:"csv_file,attr"`
			} `xml:"dimension"`
		} `xml:"dimensions"`
	}

	ProtocolConfig struct {
		Protocol     `xml:""`
		ProtocolPath string
	}
)

func NewProtocolConfig(protocolPath string, protoName string) *ProtocolConfig {
	file, err := os.Open(protocolPath + string(os.PathSeparator) + protoName + ".xml")
	if err != nil {
		log.Fatal(err)
	}
	var p = ProtocolConfig{
		ProtocolPath: protocolPath,
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err = xml.Unmarshal(data, &p); err != nil {
		log.Error(err)
	}
	return &p
}
