package main
/* 
import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	
)

var db *sql.DB
var session *mgo.Session
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(192.168.1.112:3306)/?charset=utf8mb4")
	db.Ping()
	session,_=mgo.Dial(`mongodb://dongxincbms:bms2016!@192.168.1.112:27017/admin`)

}
func main() {

	db.Query("drop database if exists tmpdb")
	db.Query("create database tmpdb")
	fmt.Println("插入数据")
	db.Query("create table tmpdb.tmptab(c1 int,c2 varchar(20),c3 varchar(20)) default charset=utf8mb4")
	_, err := db.Query("insert into tmpdb.tmptab values(101,'姓名1','address2'),(102,'abc','address2')")
	if err != nil {
		log.Fatalln(err)
	}
	qurey, err := db.Query("select * from tmpdb.tmptab")
	if err != nil {
		log.Fatalln(err)
	}
	v := reflect.TypeOf(qurey)
	fmt.Println(v)
	printResult(qurey)
	fmt.Println("插入数据2")
	stmt, err := db.Prepare(`INSERT INTO tmpdb.tmptab(c1,c2,c3) VALUES(?,?,?)`)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := stmt.Exec("103", "姓名2", "beijing")
	if err != nil {
		log.Fatalln(err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	println(num)
	qurey, err = db.Query("select * from tmpdb.tmptab")
	if err != nil {
		log.Fatalln(err)
	}
	printResult(qurey)
	r := mux.NewRouter()
	r.HandleFunc("/hellow", func(w http.ResponseWriter, r *http.Request) {
		cacheCout,err:=session.DB("cbms").C("cache").Count()
		if err!=nil{
			log.Fatalln(err)
		}
		fmt.Fprintf(w, "cache count %d",cacheCout)
	}).Methods("GET")
	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
func printResult(qurey *sql.Rows) {
	columns, _ := qurey.Columns()
	for i, c := range columns {
		fmt.Println(i, c)
	}

	values := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string)
	i := 0
	for qurey.Next() {
		if err := qurey.Scan(scans...); err != nil {
			log.Fatalln(err)
		}
		row := make(map[string]string)
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}
	for k, v := range results {
		fmt.Println(k, v)
	}
}
 */