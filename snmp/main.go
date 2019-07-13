package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/soniah/gosnmp"
)

func main() {
	gosnmp.Default.Target = "192.168.2.133"
	err := gosnmp.Default.Connect()
	if err != nil {
		log.Fatal(err)
	}
	a := [4]int{1,2,3,4}
	a1 := a[:3]
	//var b [10]int
	b1 := []int{2}
	copy(b1,a1)
	s:=[]rune{'中','国','a','b','c'}
	for _,v:=range s{
		fmt.Printf("%b",v)
		fmt.Println("=-------------")
	}
	fmt.Println(b1)
	defer gosnmp.Default.Conn.Close()
	oids := []string{".1.3.6.1.2.1.25.1", "1.3.6.1.2.1.1.7.0"}
	result, err2 := gosnmp.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}
	r,err:=gosnmp.Default.WalkAll(".1.3.6.1.4.1.2021.9.1.7")
	for _,i :=range r{
		fmt.Println(i.Name,i.Type,i.Value)
	}
	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}

	}
}
