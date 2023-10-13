#!/bin/bash

# Must use zkGo to build
GOOS=wasip1 GOARCH=wasm go build -o fib.wasm fib.go 

# Require node > 20
node ../zkWasm-emulator/wasi/wasi_exec_node.js fib.wasm /dev/null
