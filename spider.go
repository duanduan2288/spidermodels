package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/duanduan2288/spidermodels/models"
	// spider "github.com/duanduan2288/spidermodels/models/spider"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
	err := models.RedisModel.SetContent("dddd", "duan")
	fmt.Println(err)
}

func main() {
	models.RedisInit()
	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Spider 抓取
func Spider(body string) (string, error) {
	//urls := []string{}
	URL := ""
	req, err := http.NewRequest("GET", URL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	// req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respData), nil
}
