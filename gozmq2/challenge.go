package main

/* import (
	"image"
	"os"
	"image/jpeg"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
	"encoding/xml"
)

func init() {
	log.SetFlags(log.LUTC | log.Ldate | log.Lshortfile)

}

const (
	ADDR = `challenge.yuansuan.cn:7042`
)

var (
	response string
)
var all []byte

func main() {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(conn)

	msg, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf(msg)

	ss := strings.Split(msg, ":")
	s := string([]byte(ss[1])[:len(ss[1])-1])
	fmt.Fprintf(conn, "IAM:%s:a@bc.com\n", s)

	msg, _ = reader.ReadString('\n')
	fmt.Printf(msg)
	for {
		var seq = make([]byte, 4)
		n, err := reader.Read(seq)
		if err != nil {
			log.Fatalln(err)
		}
		if n != 4 {
			fmt.Println("done", n, 1)
			break
		}
		fmt.Println(seq)

		var check = make([]byte, 4)
		n, err = reader.Read(check)
		if err != nil {
			log.Fatalln(err)
		}
		if n != 4 {
			fmt.Println("done", n, 2)
			break
		}
		fmt.Println(check)

		var length = make([]byte, 4)
		n, err = reader.Read(length)

		if err != nil {
			log.Fatalln(err)
		}
		if n != 4 {
			fmt.Println("done", n, 3)
			break
		}
		l := binary.BigEndian.Uint32(length)
		if l == 0 {
			fmt.Println("done", n, 1)
			break
		}
		fmt.Println(l)

		var data = make([]byte, l)
		n, err = reader.Read(data)
		if err != nil {
			log.Fatalln(err)
		}
		fill(seq, check, data, uint32(len(data)))
		if uint32(n) != l {
			fmt.Println("done", n, 5)
			break
		}
	}
	fmt.Println(all)
	j,err:=os.Create("a.jpg")
	b:=bytes.NewBuffer(all)
	m,_,_:=image.Decode(b)
	jpeg.Encode(j,m , nil)

}

func fill(seq, check, data []byte, l uint32) {
	var ll uint32
	ll = l % 4
	if ll != 0 {
		for i := uint32(0); i < 4-ll; i++ {
			data = append(data, 0xAB)
		}
		l += 4 - ll
	}

	if checksum(seq, check, data, l) != 0 {
		fmt.Println("?????")
	} else {

		all = append(all, data...)
	}
	fmt.Println()

}

func checksum(seq, check, data []byte, l uint32) int {
	var total []byte
	total = append(total, seq...)
	total = append(total, data...)
	var a, b, sum = make([]byte, 4), make([]byte, 4), make([]byte, 4)

	a = total[0:4]
	b = total[4:8]
	for i := 0; i < 4; i++ {
		sum[i] = a[i] ^ b[i]
	}

	for i := uint32(8); i <= l; i += 4 {
		a = total[i : i+4]
		for k := 0; k < 4; k++ {
			sum[k] = a[k] ^ sum[k]
		}
	}
	fmt.Println(sum, check)
	return bytes.Compare(sum, check)
}
*/
