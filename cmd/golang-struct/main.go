// cmd/golang-struct/main.go
package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A	int8
	B	int
	C	int8
	D	int
	E	int8
}

type User2 struct {
	A	int8
	C	int8
	E	int8
	B	int
	D	int
}

func main() {
	user := User{1,2,3,4,5}
	user2 := User2{1,2,3,4,5}

	fmt.Println("User size:", unsafe.Sizeof(user))
	fmt.Println("User2 size", unsafe.Sizeof(user2))
}