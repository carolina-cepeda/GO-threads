package main

import (
	"GO-THREADS/validator"
	"fmt"
)

func main() {
	result := validator.CheckHost("200.24.34.55")
	fmt.Println("The host was found in the following blacklists:", result)
}
