package main

import "fmt"

func main() {
	var base, h uint
	fmt.Scan(&base)
	fmt.Scan(&h)
	var area float32 = float32(base*h) / 2.0
	fmt.Println(area)
}
