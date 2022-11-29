package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//urlSuffix = "_only_update_url"
	urlSuffix = ".only_update_url"
)

var tokenCheck = regexp.MustCompile(`^[a-zA-Z0-9]{32}$`)

// ServerSignConfig 静默签署扩展配置
type ServerSignConfig struct {
	SaasTemplateId    string
	ChannelTemplateId string
}

func testMap(formMap map[string]string) {
	formMap["test"] = "test"
	delete(formMap, "1")
}

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

	fmt.Println(tokenCheck.MatchString("123234"))
	fmt.Println(tokenCheck.MatchString("yDxjOUUgydjf7zv0UuO4zjEC0AKihrfi"))

	sc := ServerSignConfig{
		SaasTemplateId:    "yDRsAUUgyg1uidv9UESZxgj8Co9rytBE",
		ChannelTemplateId: "yDRsoUUgyg1xmwjuUupeCi0yCs4MwqZL",
	}

	bytes, _ := json.Marshal(sc)
	fmt.Println(string(bytes))

	tMap := make(map[string]string)
	for i := 0; i < 100; i++ {
		tMap[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	testMap(tMap)

	fmt.Println(tMap)

}
