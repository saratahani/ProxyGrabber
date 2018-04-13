package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("prox.txt", os.O_APPEND|os.O_CREATE, 0644)

	for _, v := range unique(getTag()) {
		s, _ := cleaner(v)
		f.WriteString(s)
	}
}
