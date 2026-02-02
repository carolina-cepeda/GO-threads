package workers

import "GO-THREADS/datasource"

type Result struct {
	Server int
	Found  bool
}

func SearchRange(
	dataSrc *datasource.BlacklistDataSource,
	ip string,
	start int,
	end int,
	results chan<- Result,
) {
	for i := start; i < end; i++ {
		if dataSrc.IsInBlacklist(i, ip) {
			results <- Result{Server: i, Found: true}
		}
	}
}
