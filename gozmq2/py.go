package main

/*
import (
	"fmt"

	"github.com/sbinet/go-python"
)

type Test struct {
	Str2Py func(string) *python.PyObject
	Py2Str func(*python.PyObject) string
	Int2Py func(float64) *python.PyObject
	Py2Int func(*python.PyObject) float64
}
func init()  {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}

}
func NewTest() *Test {
	test := &Test{
		Str2Py: python.PyString_FromString,
		Py2Str: python.PyString_AsString,
		Int2Py: python.PyFloat_FromDouble,
		Py2Int: python.PyFloat_AsDouble,
	}
	return test
}
func (test *Test) ImportModule(dir, name string) *python.PyObject {
	module := python.PyImport_ImportModule("sys")
	path := module.GetAttrString("path")
	python.PyList_Insert(path, 0, test.Str2Py(dir))
	return python.PyImport_ImportModule(name)
}
func main() {
	test := NewTest()
	// module := test.ImportModule("/home/chauncey/BMS/cbms/dataprovider/protocol/lang", "__init__")

	module := test.ImportModule(".", "test")
	// f := module.GetAttrString("test2")
	// argv := python.PyTuple_New(1)
	// python.PyTuple_SetItem(argv, 0, test.Int2Py(222))

	// res := f.Call(argv, python.Py_None)
	// fmt.Println(test.Py2Int(res))

	v := python.PyDict_New()
	t:=python.PyDict_New()
	python.PyDict_SetItem(t,test.Str2Py("1"),test.Int2Py(1))
	python.PyDict_SetItem(v, test.Str2Py("c"), t)

	python.PyDict_SetItem(v, test.Str2Py("a"), test.Int2Py(15262847))
	python.PyDict_SetItem(v, test.Str2Py("b"), test.Str2Py("-----222---"))
	python.PyDict_SetItem(v, test.Str2Py("FlowId"), test.Str2Py("720894889111978263:-7633500475215577087:0"))
	// err:=python.PyDict_SetItem(v, test.Str2Py("c"), python.PyString_FromString("222x"))
	// if err!=nil{
	// 	log.Fatalln(err)
	// }
	// fmt.Printf("%t\n", python.PyDict_Check(v))
	// fmt.Printf("%s\n", test.Py2Str(python.PyDict_GetItemString(v, "b")))
	r := module.GetAttrString("py_eval_expr")

	origList := python.PyTuple_New(2) //使用元组传递参数
	python.PyTuple_SetItem(origList, 0, test.Str2Py(`"WUXI" if FlowId[3:6] == "195" else "SHANGHAI"`))
	python.PyTuple_SetItem(origList, 1, v)

	res := r.Call(origList, python.Py_None)
	fmt.Println(Py2Go(res))

	origList = python.PyTuple_New(2) //使用元组传递参数
	python.PyTuple_SetItem(origList, 0, test.Str2Py(`FlowId[11:21]`))
	python.PyTuple_SetItem(origList, 1, v)

	res = r.Call(origList, python.Py_None)

	exc, val, tb := python.PyErr_Fetch() //捕获异常
	fmt.Printf(" exc=%v\n val=%v\n tb=%v\n\n", exc, val, tb)

	fmt.Println(Py2Go(res))
	// fmt.Println(Py2Go(v))
	x  :=map[int][]string {
		1:[]string{"2"},
	}
	// x=make(map[int][]string)
	// x[1]=[]string{"1"}
	x[2]=[]string{"2","3x"}
	for i,v:=range x{
		fmt.Println(i,v)
	}
	fmt.Println(1526538113/5*5)
}
type CompositeType struct{
	DictMap map[interface{}]interface{}
	ListSlice []interface{}
}
func Py2Go(res *python.PyObject) interface{} {
	var any interface{}
	if python.PyFloat_Check(res){
		any=python.PyFloat_AsDouble(res)
	}else if python.PyString_Check(res){
		any=python.PyString_AsString(res)
	}else if python.PyDict_Check(res){
		// for i:=0;i<python.PyDict_Size(res);i++{
		// 	python.PyList_GetItem(python.PyDict_Items(res), i)
		// }
		// TODO
		panic("Dict Not implemented. ")
	}

	return any
} */
/* package main

import (
	"fmt"
	"log"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		log.Panic(err)
	}
}

func main() {

	origList := python.PyList_New(0)
	python.PyList_Append(origList, python.PyString_FromString("i want this gone"))
	fmt.Println(python.PyString_AsString(origList.Str()))

	python.PyList_SetSlice(origList, 0, 1, nil)
	fmt.Println(python.PyString_AsString(origList.Str()))
} */
