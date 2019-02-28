package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
//
//var dbName = "choujiang.cb"
//var f *os.File
//
//func Exists(path string) bool {
//	_, err := os.Stat(path) //os.Stat获取文件信息
//	if err != nil {
//		if os.IsExist(err) {
//			return true
//		}
//		return false
//	}
//	return true
//}
//
//func hello(c echo.Context) (err error) {
//	if !Exists(dbName) {
//		f, err = os.Create(dbName)
//	}
//
//	result := c.FormValue("result")
//	b, err := json.Marshal(result)
//	_, err = f.Write(b)
//	m := make(map[string]interface{})
//	m["status"] = err
//	return c.JSON(http.StatusOK, m)
//}
//func hello2(c echo.Context) (err error) {
//	m := make(map[string]interface{})
//	m["result"] = err
//	 if !Exists(dbName){
//		 return c.JSON(http.StatusInternalServerError, m)
//	 }
//	f,err = os.Open(dbName)
//	b,err:=ioutil.ReadAll(f)
//	j:=make(map[string]interface{})
//	m["result"] = j
//	json.Unmarshal(b,&j)
//	return c.JSON(http.StatusOK, j)
//}
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	//e.GET("/get_cj", hello2)
	//e.GET("/cjjg", hello)
	e.File("/", "index.html")
	e.Static("/", ".")
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
