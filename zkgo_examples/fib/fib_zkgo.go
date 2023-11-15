package main

//go:wasmimport env wasm_input
//go:noescape
func wasm_input(isPublic uint32) uint64

//go:wasmimport env require
//go:noescape
func require(uint32)

func main() {
	var a0, a1 uint64
	a0 = wasm_input(1)
	a1 = wasm_input(1)
	for i := 2; i <= 1000; i++ {
		a0, a1 = a1, a0 + a1
	}
	an := wasm_input(1)
	if an != a1 {
		require(0)
	}
}
