package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	println("Listening..")
	server()
}
