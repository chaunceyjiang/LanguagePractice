package main

import (
	"flag"

	"github.com/golang/snappy"
	"io/ioutil"
	"log"
	"os"
)

var (
	input  = flag.String("i", "", "input filename")
	output = flag.String("o", "", "output filename")
	decode = flag.Bool("d", false, "decode")
	encode = flag.Bool("e", false, "encode")
)

func main() {
	flag.Parse()
	if *decode == *encode {
		log.Fatalln("请选择一个解压或者压缩")
		flag.PrintDefaults()
	}
	if *input == *output  {
		log.Fatalln("请选择一个输入文件和一个输出文件")
	}
	file, err := os.Open(*input)
	if err != nil {
		log.Fatalln(err)
	}

	in, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	out := []byte(nil)
	if *decode {
		out, err = snappy.Decode(nil, in)
		if err!=nil{
			log.Fatalln(err)
		}
		err = write2file(out)
	} else {
		out = snappy.Encode(nil,in)
		err = write2file(out)
	}
	if err!=nil{
		log.Fatalln(err)
	}

}
func write2file(out []byte) error {
	szfile, err := os.OpenFile(*output, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	_, err = szfile.Write(out)
	return err
}
