package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

var innerOpenIds = []string{
	"obLsS5f0kCahMnUEaeQwQKM0N0nY",
	"obLsS5TWNBM482WtEE7bQDSVoDdk",
	"obLsS5ZNHsb1eCWF15s0_kEpmJHk",
	"obLsS5cEfDus2KKZ-jT7MXhG5hL0",
	"obLsS5Z6csqUzYONu5_UwftslmBs",
	"obLsS5YD_NuNRxohmwLukKB1Gygw",
	"obLsS5e0rjlGMr6DRZxUlvn6ouKw",
	"obLsS5fKfMmUHgEISG-UpBsPnfGc",
	"obLsS5Tf4x_SK7F5WQWIxUHaW8gQ",
	"obLsS5VH17IuCuJSGlCwfTySoguA",
	"obLsS5bZM3uV5MuKdtqhdA4MGLt0",
	"obLsS5c8zsoFeSyIbM7kpIQPogKE",
	"obLsS5R1EAxHMsQ_bfntMv-vB9xo",
	"obLsS5bHNfpVMTePFM5iYK0WWLlY",
	"obLsS5dDXOO410Gi5_P_mWI0Dj0M",
	"obLsS5dGWr4gloZZwxZxVw5WQM3U",
	"obLsS5XMbrrJ_nygl4TaNUd4xW64",
	"obLsS5ZSbBeUCBxH4K6vlidUfgp4",
	"obLsS5eK1KmLeqZkowbDQ442kxtc",
	"obLsS5UShZKK6_PAya0rOIB4BBAg",
	"obLsS5b2TI7uazo6fTt7RjPVBHco",
	"obLsS5Zdh5EwTH1RuKXBvOlz1V3E",
	"obLsS5YDjVx8es6ocxQ8KcmAJN2A",
	"obLsS5XmCNfai5wNJk_NQT9UyNGA",
	"obLsS5QvEDizPyZwXAqv3kRAFm40",
	"obLsS5SeeD2BPY0OTAopIa5AfRIY",
	"obLsS5b0zO3NAT9RD0KTNVU-S6no",
	"obLsS5VWNqaBFC_8-gmZePV-GlQg",
	"obLsS5Uh5OsnfG3yX9JBzOS_WV-U",
	"obLsS5bZ5otYA1o1Aw0qAzip0K8s",
	"obLsS5XJW9fhj2lr-eLHeTn-SPMs",
	"obLsS5U6hNeSMdqn27j4uG-6iwbU",
	"obLsS5RWQOOxGCRatXhT0nLldONM",
	"obLsS5XFyjW5dmk119BVhIqamucU",
	"obLsS5QfT_4W28lypAs-ygZ0UIM4",
	"obLsS5Z4A2iynnI42CtnmrrbJges",
	"obLsS5Qpj0Eidgzq7q25gdwppwlQ",
	"obLsS5d2N6YVVUSN6tI-W7lb7L-E",
	"obLsS5bQQDJbE99qWyL1Sw331Goc",
	"obLsS5W8buVAkGTngRNWZ1VfZPOs",
	"obLsS5SA_89i3MzQdo7MazrD4Dxs",
	"obLsS5TUdKmStkxNQPwHW0tYuLTM",
	"obLsS5ZNrX5GfbjXf3TGyVpHfuLQ",
	"obLsS5RYh16kNmeldzmvmjKd5tOg",
	"obLsS5duIz1ft4ynitVMJ0I-zym0",
	"obLsS5dIgbhzQMDkM1oFs87p2Jm8",
	"obLsS5Quhs7x5OmL6-lovICYejXo",
	"obLsS5T98UgHbAc5HwgAf6S3XmXw",
	"obLsS5T7Jay41fRWA1A92BKpeUMM",
	"obLsS5R637qkd5ud_QS3QHSZyZL8",
	"obLsS5R3t7qyZEAPGMcuDEiQGYvM",
	"obLsS5XH45xhCInjusPEu2xCQ_Qo",
	"obLsS5antxWnV-rhf5XIqBmkoiY0",
}

func main() {
	fmt.Println(GetSourceInfo(GetterFunc(func(key string) string {
		return "get source from inner func" + " " + key
	}), ""))

	fmt.Println(GetSourceInfo(GetterFunc(GetInfoFromRedis), "key"))

	fmt.Println(GetSourceInfo(&myDb{}, "db"))

	dbs := []*myDb{
		{
			otherFiled: "1",
		},
		{
			otherFiled: "2",
		},
		{
			otherFiled: "3",
		},
	}

	for _, db := range dbs {
		curDb := db
		go func() {
			fmt.Println(curDb.otherFiled)
		}()
	}

	openIds := strings.Join(innerOpenIds, "','")
	sql := `select u.name as name, u.mobile as mobile from users as u inner join open_users as o on
	(u.verified_on != 0 and u.id = o.user_id and u.is_deleted = 0 and o.is_deleted = 0 and o.
channel = 'WEIXINAPP' and o.open_id in ('%s')) order by u.created_on desc limit 49;`
	sql = fmt.Sprintf(sql, openIds)

	fmt.Println(sql)

	time.Sleep(time.Second * 3)
	mydb := GetDbStruct()
	if mydb == nil {
		fmt.Println("mydb is nil")
	}
	//	fmt.Println(If(ans, "", mydb.otherFiled).(string))

	testMarshal()
}

func GetInfoFromRedis(key string) string {
	return "get source from redis" + " " + key
}

func GetDbStruct() *myDb {
	return nil
}

type myDb struct {
	db         *sql.DB
	otherFiled string
}

func (db *myDb) Get(key string) string {
	return "get source from myDb" + " " + key
}

// If 三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// Statement 问卷段
type Statement struct {
	Title      string   `json:"title"`      //问卷段标题
	IsRequired bool     `json:"isRequired"` // 是否必填
	Fields     []*Field `json:"fields"`
}

// Field 问卷属性
type Field struct {
	ComponentIds   []string   `json:"componentIds"` // Field 键值
	FillKey        string     `json:"fillKey"`
	FrontComponent string     `json:"frontComponent"` // 控件类型
	Label          string     `json:"label"`          //  Label 名称
	Rules          []*Rule    `json:"rules"`          //
	SensitiveWords bool       `json:"sensitiveWords"` // 是否做敏感词校验
	Placeholder    string     `json:"placeholder"`    // 提示
	Range          RangeGroup `json:"range"`          // 控件选项
}

type RangeGroup []AnyInterface

type AnyInterface interface {
}

type (

	// Rule 规则
	Rule struct {
		Required bool   `json:"required"`
		Msg      string `json:"msg"`
	}

	// Range 范围
	Range struct {
		Text  *string `json:"text"`
		Value string  `json:"value"`
	}

	// Cascade 级联数据
	Cascade struct {
		Children []*Cascade `json:"Children"`
		Name     string     `json:"name"`
		Label    string     `json:"label"`
	}
)

func (c *Cascade) String() {
}

var Cascades = []*Cascade{
	{
		Name:  "计算机",
		Label: "jisuanji",
		Children: []*Cascade{
			{
				Name:  "计算机",
				Label: "jisuanji",
			},
		},
	},
}

func testMarshal() {
	s1 := "武汉"
	//s2 := "北京"

	var f, f1 AnyInterface
	fs := make([]AnyInterface, 0, 2)
	f = &Range{
		Text:  &s1,
		Value: "wuhan",
	}

	f1 = &Range{
		//Text:  &s2,
		Value: "beijing",
	}

	fs = append(fs, f)
	fs = append(fs, f1)

	//fs1 := make([]RangeInterface, 0, 2)
	var c, c1 AnyInterface
	c = &Cascade{
		Name:  "计算机",
		Label: "jisuanji",
		Children: []*Cascade{
			{
				Name:  "软件",
				Label: "jisuanji",
				//Children: []*Cascade{},
			},
		},
	}
	c1 = &Cascade{
		Name:  "医学",
		Label: "jisuanji",
		Children: []*Cascade{
			{
				Name:  "内科",
				Label: "jisuanji",
				//Children: []*Cascade{},
			},
		},
	}

	fs = append(fs, c)
	fs = append(fs, c1)

	forms := []*Statement{
		{
			Title:      "1",
			IsRequired: false,
			Fields: []*Field{
				{
					ComponentIds: []string{
						"1",
						"2",
					},
					Range: RangeGroup(fs),
				},
			},
		},
		{
			Title:      "2",
			IsRequired: false,
			Fields: []*Field{
				{
					ComponentIds: []string{
						"111",
						"222",
					},
					Range: RangeGroup(fs),
				},
			},
		},
	}

	data, err := json.Marshal(forms)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
