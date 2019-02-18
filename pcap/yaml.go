package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


func main() {
	//m := make(map[interface{}]interface{})
	file, err := os.Open("proto.yaml")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	m:=Proto{}
	err = yaml.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

type Proto struct {
	Basic struct{
		Display_Name string`yaml:"display_name"`
		Name string
	}
	OutPut struct{
		Is_success []string`yaml:"is_success"`
	}
	Trance_Disp struct{
		PAN string`yaml:"PAN"`
	}
	Trace struct{
		Fields []string`yaml:"fields"`
		Link_field_tstid string`yaml:"link_field_tstid"`

	}
}
