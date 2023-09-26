#!/bin/bash

# Must use zkGo to build
GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go 

# Require node > 20
node ../zkWasm-emulator/wasi/wasi_exec_node.js main.wasm /dev/null
