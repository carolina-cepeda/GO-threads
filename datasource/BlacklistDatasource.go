package datasource

import "sync"

type BlacklistDataSource struct {
	blacklists map[int]map[string]bool
	mutex         sync.Mutex
}

func NewBlacklistDataSource() *BlacklistDataSource {
	dataSrc := &BlacklistDataSource{
		blacklists: make(map[int]map[string]bool),
	}

	dataSrc.blacklists[23] = map[string]bool{"200.24.34.55": true}
	dataSrc.blacklists[50] = map[string]bool{"200.24.34.55": true}
	dataSrc.blacklists[200] = map[string]bool{"200.24.34.55": true}
	dataSrc.blacklists[1000] = map[string]bool{"200.24.34.55": true}
	dataSrc.blacklists[500] = map[string]bool{"200.24.34.55": true}

	return dataSrc
}

func (dataSrc *BlacklistDataSource) GetRegisteredServersCount() int {
	return 80000
}

func (dataSrc *BlacklistDataSource) IsInBlacklist(server int, ip string) bool {
	dataSrc.mutex .Lock()
	defer dataSrc.mutex .Unlock()

	if list, ok := dataSrc.blacklists[server]; ok {
		return list[ip]
	}
	return false
}
