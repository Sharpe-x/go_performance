package main

import (
	"database/sql"
	"fmt"
)

type Getter interface {
	Get(key string) string
}

type GetterFunc func(key string) string

func (f GetterFunc) Get(key string) string {
	return f(key)
}

func GetSourceInfo(getter Getter, key string) string {
	return getter.Get(key)
}

func main() {
	fmt.Println(GetSourceInfo(GetterFunc(func(key string) string {
		return "get source from inner func" + " " + key
	}), ""))

	fmt.Println(GetSourceInfo(GetterFunc(GetInfoFromRedis), "key"))

	fmt.Println(GetSourceInfo(&myDb{}, "db"))

}

func GetInfoFromRedis(key string) string {
	return "get source from redis" + " " + key
}

type myDb struct {
	db         *sql.DB
	otherFiled string
}

func (db *myDb) Get(key string) string {
	return "get source from myDb" + " " + key
}
