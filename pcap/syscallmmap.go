package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.LUTC)
}

func main() {
	var (
		ts, index uint32
	)
	eth := flag.String("i", "", "抓取的网卡")
	debug := flag.Bool("debug", false, "打印包的长度 默认:false")
	//count := flag.Int("c", -1, "读取的次数，-1 代表死循环 默认：-1")
	flag.Parse()
	if *eth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	recv_path := fmt.Sprintf("//home/chauncey/Desktop/%s/read", *eth)
	send_path := fmt.Sprintf("/home/chauncey/Desktop/%s/write", *eth)
	data_path := fmt.Sprintf("/home/chauncey/Desktop/%s/data", *eth)
	if !Exists(recv_path, send_path, data_path) {
		log.Println(*eth, "mmap不存在")
		os.Exit(1)
	}
	file, size := openfile(send_path)
	m, err := mMap(file, size)
	if err != nil {
		log.Fatalln(err)
	}

	index = byte2uint32(m[:4])

	dataFile, size := openfile(data_path)
	dm, err := mMap(dataFile, size)
	if err != nil {
		log.Fatalln(err)
	}
	field := []string{"ts", "ns", "cap", "packetLen"}
	for i := uint32(0); i < 4; i++ {
		ts = byte2uint32(dm[index+i*4 : index+4*i+4])
		if *debug {
			fmt.Println(field[i], ts)
		}

	}
	//pack := make([]byte, ts)
	//copy(pack,dm[index+16:])
	if uint32(len(dm[index+16:])) < ts {
		fmt.Println(hex.Dump(dm[index+16:]))
	} else {
		fmt.Println(hex.Dump(dm[index+16 : index+16+ts]))
	}

}

func byte2uint32(m []byte) uint32 {
	var index uint32
	buf := bytes.NewBuffer(m)
	binary.Read(buf, binary.LittleEndian, &index)
	return index
}

func mMap(file *os.File, size int64) ([]byte, error) {
	return syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)

}

func openfile(filename string) (*os.File, int64) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	size := info.Size()
	return file, size
}

func Exists(paths ...string) bool {
	var exist bool
	for _, path := range paths {
		log.Println(path)
		_, err := os.Stat(path) //os.Stat获取文件信息
		if err != nil {
			if os.IsExist(err) {
				exist = true
				continue
			}
			return false
		}
		exist = true
	}
	return exist
}
