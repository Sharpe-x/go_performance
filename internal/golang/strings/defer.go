package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Item struct {
	Name string
}

func (i *Item) print() {
	fmt.Println(i)
}

func main() {

	/*item := new(Item)

	defer deferFunc(item)
	defer item.print()

	item = &Item{Name: "hello"}

	fmt.Println(strings.Join([]string{"a", "b", "c"}, "|"))
	fmt.Println(url.QueryEscape("劳动合同"))*/

	receiptIdToOpenIds := make(map[int][]string)
	for i := 0; i < 1000; i++ {

		if openIds, ok := receiptIdToOpenIds[i%10]; ok {
			openIds = append(openIds, strconv.Itoa(i))
		} else {
			receiptIdToOpenIds[i] = []string{strconv.Itoa(i)}
		}
	}

	fmt.Println(receiptIdToOpenIds[0])

	var item Item
	err := json.Unmarshal([]byte(""), &item)
	fmt.Println(err)
}

func deferFunc(item *Item) {
	fmt.Println(item)
}
