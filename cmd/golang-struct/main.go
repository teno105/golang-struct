// cmd/golang-struct/main.go
package main

import "fmt"

type House struct {
	Address string
	Size int
	Price float64
	Type string
}

func main() {
	var house1 House	// 기본값
	var house2 House = House{"서울시 강동구...", 28, 9.80, "아파트"}
	var house3 House = House{
		"서울시 강동구...",
		28,
		9.80,
		"아파트",
	}
	var house4 House = House{ Size: 28, Type: "아파트" }
	var house5 House = House{
		Size: 28,
		Type: "아파트",
	}

	fmt.Printf("house1 크기: %d평\n", house1.Size)
	fmt.Printf("house2 크기: %d평\n", house2.Size)
	fmt.Printf("house3 크기: %d평\n", house3.Size)
	fmt.Printf("house4 크기: %d평\n", house4.Size)
	fmt.Printf("house5 크기: %d평\n", house5.Size)
}