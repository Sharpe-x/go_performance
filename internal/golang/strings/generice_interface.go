package main

import "fmt"

// DataProcessor 泛型接口
type DataProcessor[T any] interface {
	Process(oriData T) (newData T)
	Save(data T) error
}

type DataProcessor2[T any] interface {
	int | ~struct{ Data interface{} }

	Process(data T) (newData T)
	Save(data T) error
}

type CSVProcessor struct {
}

func (c CSVProcessor) Process(oriData string) (newData string) {
	return oriData + "_CSVProcessor"
}

func (c CSVProcessor) Save(oriData string) error {
	return nil
}

func main() {
	var processor DataProcessor[string] = CSVProcessor{}
	fmt.Println(processor.Process("hello"))
	fmt.Println(processor.Save("hello"))
}
