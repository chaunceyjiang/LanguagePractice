package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// //GPS 输入样例
// const GPS = `$GPRMC,024813.640,A,3158.4608,N,11848.3737,E,10.05,324.27,150706,,,A*50`

// func main() {
// 	if strings.HasPrefix(GPS, "$GPRMC,") {
// 		index := strings.Index(GPS, "*")
// 		f := GPS[1:index]
// 		a := f[0]
// 		b := f[1]
// 		Shex := a ^ b
// 		for _, v := range f[2:] {
// 			Shex ^= uint8(v)
// 		}
// 		s := fmt.Sprintf("%x", Shex) //转换HEX
// 		S := GPS[len(GPS)-2:]
// 		if s == S {
// 			gpsf := strings.Split(GPS, ",")
// 			bjTime, location := gpsf[1], gpsf[2]
// 			if location == "A" {
// 				hh, _ := strconv.Atoi(bjTime[0:2])
// 				mm := bjTime[2:4]
// 				ss := bjTime[4:6]
// 				fmt.Printf("%s:%s:%s\n", strconv.Itoa((hh+8)%24), mm, ss)
// 			}
// 		}
// 	}

// }
