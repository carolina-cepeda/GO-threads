package validator

import (
	"GO-THREADS/datasource"
	"GO-THREADS/workers"
	"context"
	"log"
	"sync"
)

const BLACKLIST_ALARM_COUNT = 5

func CheckHost(ip string, workersCount int) []int {
	dataSrc := datasource.NewBlacklistDataSource()

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan workers.Result)
	found := []int{}
	checkedLists := 0

	total := dataSrc.GetRegisteredServersCount()
	rangeSize := total / workersCount

	var waitGroup sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		start := i * rangeSize
		end := start + rangeSize
		if i == workersCount-1 {
			end = total
		}

		waitGroup.Add(1)
		go func(s, e int) {
			defer waitGroup.Done()
			workers.SearchRange(context, dataSrc, ip, s, e, results)
		}(start, end)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()


	for results := range results {
		checkedLists += results.Checked
		if results.Found {
			found = append(found, results.Server)
			if len(found) >= BLACKLIST_ALARM_COUNT {
				cancel()
			}
		}
	}

	log.Printf("Checked Black Lists: %d of %d", checkedLists, total)

	if len(found) >= BLACKLIST_ALARM_COUNT {
		log.Printf("HOST %s Reported as NOT trustworthy", ip)
	} else {
		log.Printf("HOST %s Reported as trustworthy", ip)
	}

	return found
}

