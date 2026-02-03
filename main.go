package main

import (
	"GO-THREADS/validator"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	start := time.Now()
	
	number := 200;
	result := validator.CheckHost("200.24.34.55", number)
	fmt.Println("Blacklists:", result)

	pprof.StopCPUProfile()

	elapsed := time.Since(start)
	fmt.Println("Total execution time:", elapsed)
}
