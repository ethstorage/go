package main

//export fibonacciNth
func fibonacciNth(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

//export zkmain
func zkmain() {
	const fib10 = 55
        res := fibonacciNth(10)
	println(res)
	if fib10 != fibonacciNth(10) {
		panic(1)
	}
}

func main() {
  zkmain()
}
