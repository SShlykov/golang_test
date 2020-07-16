package main

import (
	"encoding/json"
	"fmt"
	"github.com/inancgumus/myhttp"
	"io/ioutil"
	"net/http"
	"time"
)

type resp struct {
	UserId int32  `json:"userId"`
	PostId int32  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func postHandler(addr string) {
	mh := myhttp.New(time.Second)
	response, err := mh.Get(addr)
	errHandler("cant get response from " + addr, err)

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		errHandler("Can`t read body ", err)

		var bodystr resp
		err = json.Unmarshal([]byte(body), &bodystr)
		errHandler("Can`t read json ", err)

		fmt.Printf("response: %+v\n", bodystr)
	}
}