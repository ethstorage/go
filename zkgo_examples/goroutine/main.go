package main

//go:wasmimport env wasm_input
//go:noescape
func wasm_input(isPublic uint32) uint64

//go:wasmimport env wasm_output
//go:noescape
func wasm_output(value uint64)

//go:wasmimport env require
//go:noescape
func require(uint32)


func sum(s []uint64, c chan uint64) {
	sum := uint64(0)
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []uint64{7, 2, 8, 9, 4, 0}

	c := make(chan uint64)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	wasm_output(x)
	wasm_output(y)
	wasm_output(x+y)
}
