package cache

import (
	"encoding/json"
	"os"
	"fmt"
)

var volumesCache map[string][]CacheItem = make(map[string][]CacheItem)

const volumesDirName = ".volumes"

func AddToFileCache(path string, fromVolume string, toVolume string, fromVolumeHash string, toVolumeHash string) {
	addEmptyCacheFileIfNotExists(fromVolume, toVolume)

	fileCache := GetFileCache(fromVolume, toVolume)

	// Add to cache
	cacheItem := CacheItem{path, fromVolumeHash, toVolumeHash}
	fileCache = append(fileCache, cacheItem)

	// Save to file
	fileCacheString, _ := json.MarshalIndent(fileCache, "", " ")
	saveFile(fromVolume, toVolume, fileCacheString)
}

func RemoveFromFileCache(path string, fromVolume string, toVolume string) {
	addEmptyCacheFileIfNotExists(fromVolume, toVolume)

	fileCache := GetFileCache(fromVolume, toVolume)
	var cacheItems []CacheItem
	for _, cacheItem := range fileCache {
		if cacheItem.Path == path {
			continue
		}
		cacheItems = append(cacheItems, cacheItem)
	}
	fileCacheString, _ := json.MarshalIndent(cacheItems, "", " ")
	saveFile(fromVolume, toVolume, fileCacheString)
}

func GetFileCache(fromVolume string, toVolume string) []CacheItem {
	fileCachePath := GetFileCachePath(fromVolume, toVolume)
	cacheData, err := os.ReadFile(fileCachePath)
	if err != nil {
		fmt.Println("Error occured", err)
	}

	var cacheItems []CacheItem
	_ = json.Unmarshal([]byte(cacheData), &cacheItems)
	return cacheItems
}

func GetFileCachePath(fromVolume string, toVolume string) string {
	fromVolumeHash := GenerateHash(fromVolume)
	return GetFileCacheFolder(toVolume) + fromVolumeHash + ".json"
}

func GetFileCacheFolder(toVolume string) string {
	return toVolume + "/" + volumesDirName + "/"
}

func isCacheFileExists(fromVolume string, toVolume string) bool {
	fileCachePath := GetFileCachePath(fromVolume, toVolume)
	_, err := os.Stat(fileCachePath)
	return err == nil
}

func addEmptyCacheFileIfNotExists(fromVolume string, toVolume string) {
	if !isCacheFileExists(fromVolume, toVolume) {
		saveFile(fromVolume, toVolume, []byte("{}"))
	}
}

func saveFile(fromVolume string, toVolume string, data []byte) {
	fileCachePath := GetFileCachePath(fromVolume, toVolume)
	fileCacheFolder := GetFileCacheFolder(toVolume)
	os.MkdirAll(fileCacheFolder, os.ModePerm)
	os.WriteFile(fileCachePath, data, 0644)
}
