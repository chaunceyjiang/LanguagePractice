package main
/*
import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Cell struct {
	CellX      sql.NullString
	CellY      sql.NullString
	CelleValue sql.NullFloat64
}

var mapX map[string]string
var mapY map[string]string

func init() {
	mapX = make(map[string]string)
	mapY = make(map[string]string)
	for i, v := range Index {
		mapX[v] = AsiicHandler(i + 6)
		mapY[v] = strconv.Itoa(i + 8)
	}
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(192.168.1.204:4000)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//	querySQL := "select hq.c_9 as hc_9, bh.c_9 as bc_9, sum(bh.c_11)/10000 as sum_c_11 from ( SELECT c_9,dwxtm from  C2019_000000_611_1_21B WHERE c_1=1 and left(c_6,2)=11) hq, (select c_9,c_11,dwxtm from C2019_000000_611_1_21B where c_1=2 and left(c_6,2)=11) bh where hq.dwxtm=bh.dwxtm  GROUP BY  hq.c_9,bh.c_9 ORDER BY hq.c_9 ASC, bh.c_9 ASC"
	querySQLT := `
select hq.c_9 as hc_9, bh.c_9 as bc_9, sum(bh.c_11)/10000 as sum_c_11
from 
( SELECT c_9,dwxtm from  C2019_000000_611_1_21B WHERE c_1=1 and left(c_6,2)=%s) hq, 
(select c_9,c_11,dwxtm from C2019_000000_611_1_21B where c_1=2 and left(c_6,2)=%s) bh
where hq.dwxtm=bh.dwxtm 
GROUP BY  hq.c_9,bh.c_9
ORDER BY hq.c_9 ASC, bh.c_9 ASC`
	wg := sync.WaitGroup{}
	for _, city2 := range CityCode {

		city1, _ := strconv.Atoi(os.Args[1])
		querySQL := fmt.Sprintf(querySQLT, strconv.Itoa(CityCode[city1]), strconv.Itoa(city2))
		fmt.Println(querySQL)
		wg.Add(1)
		fmt.Println(time.Now().Unix())

		Process(&wg, querySQL, db, strconv.Itoa(CityCode[city1]), strconv.Itoa(city2))
		fmt.Println(time.Now().Unix())
	}

	wg.Wait()

}

func Process(wg *sync.WaitGroup, querySQL string, db *sql.DB, city1, city2 string) {
	defer wg.Done()
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	file, err := excelize.OpenFile("/root/stats_office/template/remote1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	wgs := New(10000)
	for rows.Next() {
		var X sql.NullString
		var Y sql.NullString
		var value sql.NullFloat64
		if err = rows.Scan(&X, &Y, &value); err != nil {
			log.Warn(err)
		}

		wgs.Add(1)
		Axis(wgs, X, Y, value, file)
	}
	wgs.Wait()
	file.SetActiveSheet(1)
	err = os.MkdirAll(fmt.Sprintf("/root/stats_office/test/%s",city1),0755)
	if err!=nil{
		log.Error(err)
	}
	err = file.SaveAs(fmt.Sprintf("/root/stats_office/test/%s/%s_%s.xlsx",city1, city1, city2))
	if err != nil {
		log.Error(err)
	}

}

func Axis(wg *pool, x sql.NullString, y sql.NullString, value sql.NullFloat64, file *excelize.File) {
	defer wg.Done()
	fmt.Println("X",x)
	fmt.Println("Y",y)

	fmt.Println("V",value)

	if x.Valid && y.Valid {
		X, err := x.Value()
		if err != nil {
			log.Error(err)
		}
		Y, err := y.Value()
		if err != nil {
			log.Error(err)
		}
		V, err := value.Value()
		if err != nil {
			log.Error(err)
		}
		err = file.SetCellValue("Sheet1", checkX(X.(string))+checkY(Y.(string)), V)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("1", checkX(X.(string))+checkY(Y.(string)), V)
	} else if x.Valid {
		X, err := x.Value()
		if err != nil {
			log.Error(err)
		}
		v, err := value.Value()
		if err != nil {
			log.Error(err)
		}
		err = file.SetCellValue("Sheet1", checkX(X.(string))+"7", v)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("2", x, y, value)
	} else if y.Valid {
		Y, err := y.Value()
		if err != nil {
			log.Error(err)
		}
		v, err := value.Value()
		if err != nil {
			log.Error(err)
		}
		err = file.SetCellValue("Sheet1", "E"+checkY(Y.(string)), v)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("3", x, y, value)
	} else if !x.Valid && !y.Valid {
		v, err := value.Value()
		if err != nil {
			log.Error(err)
		}
		err = file.SetCellValue("Sheet1", "E7", v)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("4", x, y, value)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

}

func checkX(v string) string {
	if e, ok := mapX[v]; ok {
		return e
	}
	return "BAH"
}
func checkY(v string) string {
	if e, ok := mapY[v]; ok {
		return e
	}
	return "1388"
}
*/