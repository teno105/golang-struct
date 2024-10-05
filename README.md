아래는 실습 순서에 맞춰 다시 작성한 `README.md`입니다.

---

# golang-struct

`golang-struct`는 Golang으로 작성된 간단한 애플리케이션으로, struct의 사용을 익히기 위한 실습입니다.


## 프로젝트 구조

```plaintext
golang-struct/
│
├── cmd/
│   └── golang-struct/
│       └── main.go
│
├── go.mod
├── Makefile
└── README.md
```

## 실습 순서

### 1. 패키지 구조를 위한 디렉토리 생성

먼저 프로젝트 디렉터리를 설정하고 필요한 디렉터리들을 생성합니다.

```bash
mkdir golang-struct
cd golang-struct
go mod init golang-struct

mkdir -p cmd/golang-struct
```

### 2. `main.go` 생성

`cmd/golang-struct/` 디렉터리 아래에 `main.go` 파일을 생성하고,
struct 를 선언 및 활용하는 코드를 작성합니다.

```go
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
	var house House
	house.Address = "서울시 강동구..."
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입:", house.Type)
}
```

이 코드를 통해 프로그램을 실행하면 구조체의 정보가 출력됩니다.

### 3. `Makefile` 작성

이제 프로젝트의 빌드 및 실행을 자동화하기 위한 `Makefile`을 프로젝트 루트에 작성합니다.

```makefile
# Go 관련 변수 설정
APP_NAME := helloworld
CMD_DIR := ./cmd/helloworld
BUILD_DIR := ./build

.PHONY: all clean build run test fmt vet install

all: build

# 빌드 명령어
build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

# 실행 명령어
run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# 코드 포맷팅
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# 코드 분석
vet:
	@echo "==> Running go vet..."
	go vet ./...

# 의존성 설치
install:
	@echo "==> Installing dependencies..."
	go mod tidy

# 테스트 실행
test:
	@echo "==> Running tests..."
	go test -v ./...

# 빌드 정리
clean:
	@echo "==> Cleaning build directory..."
	rm -rf $(BUILD_DIR)
```

`Makefile`을 이용하여 코드를 빌드하고 실행할 수 있습니다.

```bash
make run
```

이 명령어를 통해 `main.go`에서 작성한 `Hello, World!` 메시지를 확인할 수 있습니다.

### 4. 구조체 변수 초기화

구조체 변수를 초기화 하는 방법에 대해 알아보겠습니다.

```go
// cmd/helloworld/main.go
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
```

이제 `make run` 명령을 사용하면 각 house의 크기가 출력됩니다.

```bash
make run
```