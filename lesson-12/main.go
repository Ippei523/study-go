package main

// 関数
// 無名関数

func main() {

	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	println(i)
}
