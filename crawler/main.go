package main

import (
	"bytes"
	"crawler/fake"
	"crawler/info"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {

	req, err := http.NewRequest("GET", "https://movie.douban.com/top250?start=75&filter=", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", fake.GetUserAgent())
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("CheckRedirect")
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	buf := bytes.NewBuffer(body)
	dom, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		log.Error(err)
	}
	wg := &sync.WaitGroup{}
	dom.Find(".info").Each(func(i int, selection *goquery.Selection) {
		wg.Add(1)
		go HandleMoive(wg,selection)
	})
	wg.Wait()
}

func HandleMoive(wg *sync.WaitGroup,selection *goquery.Selection) {
	defer wg.Done()
	movie := info.NewMovie()
	hd := selection.Find("div[class=hd]")
	a := hd.Find("a")
	val, exists := a.Attr("href")
	if exists {
		movie.Url = val
	}
	a.Find("span").Each(func(i int, selection *goquery.Selection) {
		movie.Title = append(movie.Title, selection.Text())
	})
	bd := selection.Find("div[class=bd]")
	movie.Described = strings.Trim(bd.Find("p:first-child").Text()," \n")
	movie.Rating = strings.Trim(bd.Find("div[class=star] span:nth-child(2)").Text()," \n")
	movie.Num = strings.Trim(bd.Find("div[class=star] span:last-of-type").Text()," \n人评价")
	movie.Quote = strings.Trim(bd.Find("p:last-child").Text()," \n")
	fmt.Println(movie)
}