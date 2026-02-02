package main

import (
	"GO-THREADS/workers"
	"sync"
)

func main() {
	c1 := workers.CountWorker{A: 0, B: 99}
	c2 := workers.CountWorker{A: 100, B: 199}
	c3 := workers.CountWorker{A: 200, B: 299}

	//var waitGroup sync.WaitGroup
	//waitGroup.Add(3)

	// concurrent
	
	//go c1.Start(&waitGroup)
	//go c2.Start(&waitGroup)
	//go c3.Start(&waitGroup)

	//waitGroup.Wait()

	// forma secuencial 
	c1.Run()
	c2.Run()
	c3.Run()
	

}
