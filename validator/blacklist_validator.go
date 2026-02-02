package validator

import (
	"GO-THREADS/datasource"
	"GO-THREADS/workers"
)

const BLACKLIST_ALARM_COUNT = 5

func CheckHost(ip string) []int {
	dataSrc := datasource.NewBlacklistDataSource()
	results := make(chan workers.Result)
	found := []int{}

	workersCount := 8
	rangeSize := dataSrc.GetRegisteredServersCount() / workersCount

	for i := 0; i < workersCount; i++ {
		start := i * rangeSize
		end := start + rangeSize
		go workers.SearchRange(dataSrc, ip, start, end, results)
	}

	for len(found) < BLACKLIST_ALARM_COUNT {
		r := <-results
		if r.Found {
			found = append(found, r.Server)
		}
	}

	return found
}
