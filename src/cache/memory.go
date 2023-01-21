package cache

func LoadToMemory(fromVolume string, toVolume string) {
	if isCacheFileExists(fromVolume, toVolume) {
		fileCache := GetFileCache(fromVolume, toVolume)
		for _, item := range fileCache {
			AddToMemoryCache(toVolume, item.Path, item.FromVolumeHash, item.ToVolumeHash)
		}
	}
}

func AddToMemoryCache(volume string, path string, fromVolumeHash string, toVolumeHash string) {
	cacheItem := CacheItem{path, fromVolumeHash, toVolumeHash}
	volumesCache[volume] = append(volumesCache[volume], cacheItem)
}

func FindInMemory(volume string, fileHash string) bool {
	items := volumesCache[volume]
	for _, item := range items {
		if item.FromVolumeHash == fileHash {
			return true
		}
	}
	return false
}
