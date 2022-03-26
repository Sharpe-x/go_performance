package main

import (
	"container/list"
	"fmt"
	"strings"
)

type ErrorMap struct {
	keys   map[string]*list.Element
	errors *list.List
}

type errEntry struct {
	last16Str  string
	lineNumber string
	errorCount int
}

func (eMap *ErrorMap) Add(entry *errEntry) {

	if elem, ok := eMap.keys[entry.last16Str+"_"+entry.lineNumber]; ok {
		e := elem.Value.(*errEntry)
		e.errorCount++
	} else {
		elem = eMap.errors.PushBack(entry)
		eMap.keys[entry.last16Str+"_"+entry.lineNumber] = elem
	}

	if eMap.errors.Len() > 8 {
		//elem := eMap.errors.Front()
		// 循环记录时，只以第一次出现的顺序为准，后面重复的不会更新它的出现时间，仍以第一次为准
		// 所以不用删除记录 keys 记录的值
		//e := elem.Value.(*errEntry)
		//delete(eMap.keys, e.last16Str+"_"+e.lineNumber)
		eMap.errors.Remove(eMap.errors.Front())
	}
}

func (eMap *ErrorMap) Print() {
	elem := eMap.errors.Front()
	for elem != nil {
		entry := elem.Value.(*errEntry)
		fmt.Printf("%s %s %d\n", entry.last16Str, entry.lineNumber, entry.errorCount)
		elem = elem.Next()
	}
}

func main() {

	errMap := &ErrorMap{
		keys:   map[string]*list.Element{},
		errors: list.New(),
	}
	var (
		errStr, line string
	)
	for {

		_, err := fmt.Scanf("%s %s\n", &errStr, &line)
		if err != nil || len(errStr) == 0 {
			break
		}

		strs := strings.Split(errStr, "\\")
		if err != nil {
			break
		}

		lastNum := len(strs[len(strs)-1])
		start := 0
		if lastNum >= 16 {
			start = lastNum - 16
		}
		errMap.Add(&errEntry{
			lineNumber: line,
			last16Str:  strs[len(strs)-1][start:],
			errorCount: 1,
		})
	}
	errMap.Print()
}
