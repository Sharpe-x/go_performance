package main

// A Getter loads data for a key.
// 定义了一个接口 Getter，只包含一个方法 Get(key string) ([]byte, error)
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
// 定义了一个函数类型 GetterFunc，GetterFunc 参数和返回值与 Getter 中 Get 方法是一致的。
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
// GetterFunc 还定义了 Get 方式，并在 Get 方法中调用自己，这样就实现了接口 Getter。所以 GetterFunc 是一个实现了接口的函数类型，简称为接口型函数。
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// 接口型函数只能应用于接口内部只定义了一个方法的情况，例如接口 Getter 内部有且只有一个方法 Get。既然只有一个方法，为什么还要多此一举，封装为一个接口呢？定义参数的时候，直接用 GetterFunc
/*这个函数类型不就好了，让用户直接传入一个函数作为参数，不更简单吗？

既能够将普通的函数类型（需类型转换）作为参数，也可以将结构体作为参数，使用更为灵活，可读性也更好，这就是接口型函数的价值。
所以呢，接口型函数的价值什么？
https://geektutu.com/post/7days-golang-q1.html
*/
