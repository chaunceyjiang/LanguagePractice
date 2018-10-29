package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//	"strings"
//	"os/user"
//)
//
////ReaderExample 样例
//func ReaderExample() {
//FOREND:
//	for {
//		readerMenu()
//		var (
//			ch   string
//			data []byte
//			err  error
//		)
//		fmt.Scanln(&ch)
//		switch strings.ToLower(ch) {
//		case "1":
//			fmt.Println("请输入不多于9个字符，以回车结束：")
//			data, err = ReadFrom(os.Stdin, 11)
//		case "2":
//			users,err := user.Current()
//			file, err := os.Open(users.HomeDir+"/a.txt")
//			if err != nil {
//				fmt.Println("打开文件失败！：错误 ", err)
//				continue
//			}
//			data, err = ReadFrom(file, 1)
//		case "b":
//			goto FOREND
//		case "q":
//			os.Exit(0)
//		default:
//			fmt.Println("输入错误！")
//			continue
//		}
//		if err != nil {
//			fmt.Println("数据读取失败！", err)
//
//		} else {
//			fmt.Printf("%s", data)
//		}
//	}
//}
//func readerMenu() {
//	fmt.Println("")
//	fmt.Println("*******从不同来源读取数据*******")
//	fmt.Println("*******请选择数据源，请输入：****")
//	fmt.Println("1 标准输入")
//	fmt.Println("2 普通文件")
//	fmt.Println("3 从字符串")
//	fmt.Println("4 从网络")
//	fmt.Println("b 返回上级菜单")
//	fmt.Println("q 退出")
//	fmt.Printf("%s", ":")
//}
//
////ReadFrom 获取数据
//func ReadFrom(reader io.Reader, num int) ([]byte, error) {
//	p := make([]byte, num)
//	n, err := reader.Read(p)
//	if n > 0 {
//		return p[:n], nil
//	}
//	return p, err
//}
//
//func main() {
//	ReaderExample()
//}
