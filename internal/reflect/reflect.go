package reflect

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

//使用反射赋值，效率非常低下，如果有替代方案，尽可能避免使用反射，特别是会被反复调用的热点代码。例如 RPC 协议中，需要对结构体进行序列化和反序列化，这个时候避免使用 Go 语言自带的 json 的 Marshal
//和 Unmarshal 方法，因为标准库中的 json 序列化和反序列化是利用反射实现的。可选的替代方案有 easyjson，在大部分场景下，相比标准库，有 5 倍左右的性能提升。

type Config struct {
	Name    string `json:"server-name"`
	Ip      string `json:"server-ip"`
	URL     string `json:"server-url"`
	TimeOut string `json:"time-out"`
}

func readConfig() *Config {

	config := Config{}

	typ := reflect.TypeOf(config)
	val := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			// 利用反射获取到 Config 的每个字段的 Tag 属性，拼接出对应的环境变量的名称。
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			// 查看该环境变量是否存在，如果存在，则将环境变量的值赋值给该字段。
			if env, exist := os.LookupEnv(key); exist {
				val.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}

	return &config
}

func StartReflect() {
	_ = os.Setenv("CONFIG_SERVER_NAME", "global_server")
	_ = os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	_ = os.Setenv("CONFIG_SERVER_URL", "http://127.0.0.1:")
	fmt.Printf("%+v", readConfig())
}
