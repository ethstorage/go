package main

import "fmt"

func main() {
	var a0, a1, an uint64;
	fmt.Scan(&a0);
	fmt.Scan(&a1);
	for i := 2; i <= 1000; i++ {
		an = a0 + a1;
		a0 = a1;
		a1 = an;
	}
	fmt.Printf("%d", an);
}
