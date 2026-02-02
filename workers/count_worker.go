package workers

import (
	"fmt"
	"sync"
)

/*
	structure of the 'object'
*/
type CountWorker struct {
	A int
	B int
}

func (c *CountWorker) Run() {
	for i := c.A; i <= c.B; i++ {
		fmt.Println(i)
	}
}
/*
	start alike of java threads
*/

func (c *CountWorker) Start(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := c.A; i <= c.B; i++ {
		fmt.Println(i)
	}
	}
