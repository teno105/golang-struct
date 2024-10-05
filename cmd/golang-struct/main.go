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

	fmt.Printf("house1 주소: %s, 크기: %d평 가격: %.2f억 원 타입: %s\n", house1.Address, house1.Size, house1.Price, house1.Type)
	fmt.Printf("house2 주소: %s, 크기: %d평 가격: %.2f억 원 타입: %s\n", house2.Address, house2.Size, house2.Price, house2.Type)
	fmt.Printf("house3 주소: %s, 크기: %d평 가격: %.2f억 원 타입: %s\n", house3.Address, house3.Size, house3.Price, house3.Type)
	fmt.Printf("house4 주소: %s, 크기: %d평 가격: %.2f억 원 타입: %s\n", house4.Address, house4.Size, house4.Price, house4.Type)
	fmt.Printf("house5 주소: %s, 크기: %d평 가격: %.2f억 원 타입: %s\n", house5.Address, house5.Size, house5.Price, house5.Type)
}