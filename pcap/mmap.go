package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/exp/mmap"
	"log"
	"os"
	"time"
)

func main() {
	var r, d *mmap.ReaderAt
	var err error
	var ts, tss uint32
	var index int64
	eth := flag.String("i", "", "抓取的网卡")
	debug := flag.Bool("debug", false, "打印包的长度 默认:false")
	count := flag.Int("c", -1, "读取的次数，-1 代表死循环 默认：-1")
	flag.Parse()
	if *eth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	recv_path := fmt.Sprintf("/dev/shm/%s_cmd_recv", *eth)
	send_path := fmt.Sprintf("/dev/shm/%s_cmd_send", *eth)
	data_path := fmt.Sprintf("/dev/shm/%s_data", *eth)
	if !Exists(recv_path, send_path, data_path) {
		log.Println(*eth, "mmap不存在")
		os.Exit(1)
	}
	timestamp := make([]byte, 4)

	r, err = mmap.Open(recv_path)
	if err != nil {
		log.Fatalln(err)
	}
	r.ReadAt(timestamp, 0)
	buf := bytes.NewBuffer(timestamp)
	binary.Read(buf, binary.LittleEndian, &tss)

	index = int64(tss)

	if *debug {
		fmt.Println("index", tss)
	}

	d, err = mmap.Open(data_path)
	if err != nil {
		log.Println(err)
	}
	field := []string{"ts", "ns", "cap", "packetLen"}
	var c int
	for {
		if c > *count && *count >= 0 {
			break
		}
		for i := int64(0); i < 4; i++ {

			d.ReadAt(timestamp, index+i*4)
			buf1 := bytes.NewBuffer(timestamp)
			binary.Read(buf1, binary.LittleEndian, &ts)
			if *debug {
				fmt.Println(field[i], ts)
			}
		}

		index += 16
		if *debug {
			fmt.Println("body_index", index)
		}

		pack := make([]byte, ts)
		d.ReadAt(pack, index)
		fmt.Println(hex.Dump(pack))
		time.Sleep(10)
		index += int64(ts)
		c++
	}

}
//func Exists(paths ...string) bool {
//	var exist bool
//	for _, path := range paths {
//		_, err := os.Stat(path) //os.Stat获取文件信息
//		if err != nil {
//			if os.IsExist(err) {
//				exist = true
//				continue
//			}
//			return false
//		}
//		exist = true
//	}
//	return exist
//}
