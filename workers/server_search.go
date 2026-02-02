package workers

import (
	"GO-THREADS/datasource"
	"context"
)

type Result struct {
	Server  int
	Found   bool
	Checked int
}


func SearchRange(
    context context.Context,
    dataSrc *datasource.BlacklistDataSource,
    ip string,
    start int,
    end int,
    results chan<- Result,
) {
    for i := start; i < end; i++ {
        select {
        case <-context.Done():
            return
        default:
        }
        
        res := Result{Checked: 1}

        if dataSrc.IsInBlacklist(i, ip) {
            select {
            case <-context.Done():
                return
            case results <- Result{Server: i, Found: true}:
            }
        }
        select {
            case <-context.Done():
            return
        case results <- res:
        }
        
    }
}
