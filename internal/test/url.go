package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println(url.QueryEscape("https://partner.ess.tencent.com/100018139744620e58a075879"))
	fmt.Println(url.QueryEscape("https://console.cloud.tencent.com/cam/capi"))
	// https%3A%2F%2Fconsole.cloud.tencent.com%2Fcam%2Fcapi
	//
}

// https%3A%2F%2Fpartner.ess.tencent.com%2F100018139744620e58a075879
// https%3A%2F%2Fpartner.ess.tencent.com%2F100018139744620e58a075879
