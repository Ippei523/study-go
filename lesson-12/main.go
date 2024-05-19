package main

// 関数
// 無名関数

func main() {

	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	println(i)

	f2, _ := func(x, y int) (int, int) {
		q := x / y
		r := x % y
		return q, r
	}(9, 4)

	println(f2)
}
