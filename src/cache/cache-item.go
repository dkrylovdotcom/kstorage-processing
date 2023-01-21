package cache

type CacheItem struct {
	Path           string `json:"path"`
	FromVolumeHash string `json:"fromVolumeHash"`
	ToVolumeHash   string `json:"toVolumeHash"`
}
