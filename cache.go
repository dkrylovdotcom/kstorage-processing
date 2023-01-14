package main

import (
	"os"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

var volumesCache map[string][]CacheItem = make(map[string][]CacheItem)

type CacheItem struct {
	Path string `json:"path"`
	FromVolumeHash string `json:"fromVolumeHash"`
	ToVolumeHash string `json:"toVolumeHash"`
}

func addToFileCache(path string, fromVolume string, toVolume string, fromVolumeHash string, toVolumeHash string) {
	if (!isVolumeJsonExists(fromVolume, toVolume)) {
		saveFile(fromVolume, toVolume, []byte ("{}"))
	}
	volumeJson := getVolumeJson(fromVolume, toVolume)

	// Add to cache
	cacheItem := CacheItem{path, fromVolumeHash, toVolumeHash}
	volumeJson = append(volumeJson, cacheItem)

	// Save to file
	volumeJsonString, _ := json.MarshalIndent(volumeJson, "", " ")
	saveFile(fromVolume, toVolume, volumeJsonString)
}

func removeFromFileCache(path string, fromVolume string, toVolume string) {
	if (!isVolumeJsonExists(fromVolume, toVolume)) {
		saveFile(fromVolume, toVolume, []byte ("{}"))
	}
	volumeJson := getVolumeJson(fromVolume, toVolume)
	var cacheItems []CacheItem
	for _, cacheItem := range volumeJson {
		if (cacheItem.Path == path) {
			continue
		}
		cacheItems = append(cacheItems, cacheItem)
	}
	volumeJsonString, _ := json.MarshalIndent(cacheItems, "", " ")
	saveFile(fromVolume, toVolume, volumeJsonString)
}

func saveFile(fromVolume string, toVolume string, data []byte) {
	volumeJsonPath := getVolumeJsonPath(fromVolume, toVolume)
	volumeJsonFolder := getVolumeJsonFolder(toVolume)
	os.MkdirAll(volumeJsonFolder, os.ModePerm)
	os.WriteFile(volumeJsonPath, data, 0644)
}

func loadToMemory(fromVolume string, toVolume string) {
	if (isVolumeJsonExists(fromVolume, toVolume)) {
		volumeJson := getVolumeJson(fromVolume, toVolume)
		for _, item := range volumeJson {
			addToMemoryCache(toVolume, item.Path, item.FromVolumeHash, item.ToVolumeHash)
		}
	}
}

func addToMemoryCache(volume string, path string, fromVolumeHash string, toVolumeHash string) {
	cacheItem := CacheItem{path, fromVolumeHash, toVolumeHash}
	volumesCache[volume] = append(volumesCache[volume], cacheItem)
}

func findInMemory(volume string, fileHash string) bool {
	items := volumesCache[volume]
	for _, item := range items {
		if (item.FromVolumeHash == fileHash) {
			return true
		}
	}
	return false
}

func generateHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
