package reflect

import (
	"reflect"
	"testing"
)

func TestReflectStart(t *testing.T) {
	StartReflect()
}

// 反射创建对象的耗时约为 new 的 1.5 倍，相差不是特别大。
func BenchmarkNew(b *testing.B) {
	var config *Config
	for i := 0; i < b.N; i++ {
		config = new(Config)
	}
	_ = config
}

func BenchmarkNewReflect(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

// =================================
func BenchmarkSet(b *testing.B) {
	config := new(Config)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		config.Name = "name"
		config.Ip = "ip"
		config.URL = "url"
		config.TimeOut = "timeout"
	}
}

func BenchmarkReflect_FieldSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("name")
		ins.Field(1).SetString("ip")
		ins.Field(2).SetString("url")
		ins.Field(3).SetString("TimeOut")
	}
}

func BenchmarkReflect_FieldByNameSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("name")
		ins.FieldByName("Ip").SetString("ip")
		ins.FieldByName("URL").SetString("url")
		ins.FieldByName("TimeOut").SetString("timeout")
	}
}

// FieldByName 相比于 Field 有一个数量级的性能劣化。那在实际的应用中，就要避免直接调用 FieldByName。
//我们可以利用字典将 Name 和 Index 的映射缓存起来。避免每次反复查找，耗费大量的时间。

func BenchmarkReflect_FieldByNameCacheSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	cache := make(map[string]int)
	for i := 0; i < typ.NumField(); i++ {
		cache[typ.Field(i).Name] = i
	}

	ins := reflect.New(typ).Elem()

	for i := 0; i < b.N; i++ {
		ins.Field(cache["Name"]).SetString("name")
		ins.Field(cache["IP"]).SetString("ip")
		ins.Field(cache["URL"]).SetString("url")
		ins.Field(cache["Timeout"]).SetString("timeout")
	}
}
