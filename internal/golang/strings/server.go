package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	// 接收消息
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("我累了,想休息,你待会再来找我吧 %v", err.Error())
	}

	fmt.Println(string(data))
	io.WriteString(w, "hello, world!\n")
}
func main() {
	http.HandleFunc("/ChatGpt/Api", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
