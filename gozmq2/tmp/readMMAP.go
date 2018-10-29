package main
/* 
import (
	"bufio"
	"fmt"
	"os"

	"github.com/fabiokung/shm"
)

func main() {
	file, err := shm.Open("eth1_cmd_recv", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer shm.Unlink(file.Name())
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadByte()
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%d", data)
	}
}
 */