package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type T struct {
	Content string `json:"content"`
}
type Msg struct {
	Msgtype string `json:"msgtype"`
	Text    T      `json:"text"`
}

func main() {
	var t string
	cmd := exec.Command("node", "-h")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t = err.Error()
	} else {
		if err = cmd.Start(); err != nil {
			t = err.Error()
		}
		b, err := ioutil.ReadAll(stdout)
		if err != nil {
			t = err.Error()
		}
		if err = cmd.Wait(); err != nil {
			t = err.Error()
		}
		fmt.Println(t)
		fmt.Println(string(b))
	}
	fmt.Println(t)
	//tokenurl := `https://oapi.dingtalk.com/robot/send?access_token=3d97d8f8d16f930f3305080594c5d6b7f73fea2989ed2b474164382e260caa45`
	//client := &http.Client{}
	//body := Msg{"text",T{t}}
	//
	//buf := new(bytes.Buffer)
	//_ = json.NewEncoder(buf).Encode(body)
	//req, err := http.NewRequest("POST", tokenurl, buf)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//req.Header.Set("Content-Type", "application/json")
	//
	//res, e := client.Do(req)
	//if e != nil {
	//	log.Fatal(e)
	//}
	//
	//defer res.Body.Close()
	//_, _ = io.Copy(os.Stdout, res.Body)
}
