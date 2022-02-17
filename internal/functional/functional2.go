package main

import "fmt"

func a(str string) string {
	fmt.Println("a_" + str)
	return "a_" + str
}

func b(str string) string {
	return "b_" + str
}

func c(str string) string {
	return "c_" + str
}

func aFunc(fn func(s string) string) func(string) string {
	fmt.Println("aFunc")
	return fn //a_word
}

func bFunc(fn func(s string) string) func(string) string {
	fmt.Println("bFunc")
	return fn
}

func cFunc(fn func(s string) string) func(string) string {
	fmt.Println("cFunc")
	return fn
}

//===================
func aaFunc(strIn string, fn func(s string) string) (strOut string) {
	fmt.Println("aaFunc start")
	strOut = fn(strIn)
	fmt.Println("aaFunc end")
	return strOut
}

/*func bbFunc(strIn string, fn func(s string) string) (strOut string) {
	fmt.Println("bbFunc start")
	strOut = fn(strIn)
	fmt.Println("bbFunc end")
	return strOut
}

func ccFunc(strIn string, fn func(s string) string) (strOut string) {
	fmt.Println("ccFunc start")
	strOut = fn(strIn)
	fmt.Println("ccFunc end")
	return strOut
}*/

func main() {
	print(a(b(c("word"))))

	fmt.Println("\n=========")
	fn := cFunc(bFunc(aFunc(a)))
	fmt.Println(fn("word"))
	fmt.Println("=========")
	//cFunc(bFunc(aFunc(a)))

	chainer := func(interceptor func(string, func(string) string) string,
		handler func(string) string) func(string) string {

		return func(strIn string) string {
			return interceptor(strIn, handler)
		}
	}

	chainerHandler := chainer(aaFunc, a)
	chainerHandler("word")

}
