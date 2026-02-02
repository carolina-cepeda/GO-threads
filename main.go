package main

import (
	"GO-THREADS/validator"
	"fmt"
)

func main() {
	result := validator.CheckHost("200.24.34.55", 8)
	fmt.Println("Blacklists:", result)
}
