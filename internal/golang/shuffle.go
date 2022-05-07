package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	//urlSuffix = "_only_update_url"
	urlSuffix = ".only_update_url"
)

func main() {
	fmt.Println("YYYY-MM-DD : ", time.Now().AddDate(0, -1, 0).Format("200601"))
	/*	if strings.HasSuffix("v1.0.0_only_update_url", "_only_update_url") {
			fmt.Println("v1.0.0")
		}
	*/
	version := "v202204291600.only_update_url_1651805507"

	// 只是更换url或者不需要重新计算banner图轮询时间或者banner图顺序不变 内部版本号不用改变
	if strings.Contains(version, urlSuffix) {
		version = strings.Split(version, urlSuffix)[0]
	}

	fmt.Println(version)
}
