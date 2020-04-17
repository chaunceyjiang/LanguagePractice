package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var filename = flag.String("f", "", "file full path")
var offset = flag.Int64("i", 0, "file seek offset")
var direction = flag.Int("d", 0, "seek direction (0,1,2)")

func init() {

}
func main() {
	flag.Parse()
	if *filename == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	var head = make([]byte, 16)
	file.Seek(*offset, *direction)
	io.ReadFull(file, head)

	l := binary.LittleEndian.Uint32(head[12:])
	fmt.Println(binary.LittleEndian.Uint32(head[0:4]))
	fmt.Println(binary.LittleEndian.Uint32(head[4:8]))
	fmt.Println(binary.LittleEndian.Uint32(head[8:12]))
	fmt.Println(binary.LittleEndian.Uint32(head[12:]))
	var body = make([]byte, l)
	io.ReadFull(file, body)
	fmt.Println(hex.Dump(body))
}
