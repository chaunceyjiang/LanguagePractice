package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)
func main()  {
	db,err:=sql.Open("mysql","root@tcp(192.168.1.204:4000)/test?charset=utf8")
	if err!=nil{
		log.Error(err)
	}
	if err = db.Ping();err!=nil{
		log.Fatal(err)
	}

	rows ,err:=db.Query(`            SELECT  UPPER(tyshxydm) as tyshxydm, b1141 as b104, c_sjclddm, zzid from c2019_000000_611_1 where b110 = 1
            union all
            SELECT b205, 1 as b110, UPPER(tyshxydm) as tyshxydm, b101, b102, b201, b0152, b0153, b0154, b0155, b0156, b0157, b0158, b2031, b2032, b2033, b2034, b103b17, b104, 0 as c_sjclddm, zzid from c2019_000000_601_1`)
	if err!=nil{
		log.Error(err)
	}
	defer rows.Close()
	for rows.Next(){
		var zzid ,bgq,code ,jldwdm sql.NullString
		if err = rows.Scan(&zzid,&bgq,&code ,&jldwdm);err!=nil{
			log.Error(err)
		}
		fmt.Println(zzid,bgq,code,jldwdm)
	}
}