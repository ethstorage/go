#!/bin/bash

# Must use zkGo to build
GOOS=wasip1 GOARCH=wasm go build -o fib.wasm fib_zkgo.go 

echo 0 | python3 write_witness.py > input.dat
echo 1 | python3 write_witness.py >> input.dat

# Require node > 20
node ../zkWasm-emulator/wasi/wasi_exec_node.js fib.wasm input.dat
