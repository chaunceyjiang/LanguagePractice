package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func LoadCsv(filename string) map[string]string {
	// 这个函数有待优化,处理每个包,都需要去重新读取文件,而该文件不会变化
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	data := make(map[string]string)
	//data, _ = r.ReadAll()
	//fmt.Println(data)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		data[line[0]] = line[1]
	}
	return data
}
