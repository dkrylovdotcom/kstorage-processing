package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

func getVolumeName(volume string) string {
	return filepath.Base(volume)
}

func isVolumeJsonExists(fromVolume string, toVolume string) bool {
	volumeJsonPath := getVolumeJsonPath(fromVolume, toVolume)
	_, err := os.Stat(volumeJsonPath);
	return err == nil
}

func getVolumeJson(fromVolume string, toVolume string) []CacheItem {
	volumeJsonPath := getVolumeJsonPath(fromVolume, toVolume)
	cacheData, err := os.ReadFile(volumeJsonPath)
	if (err != nil) {
		fmt.Println("Error occured", err)
	}

	var cacheItems []CacheItem
	_ = json.Unmarshal([]byte(cacheData), &cacheItems)
	return cacheItems
}

func getVolumeJsonPath(fromVolume string, toVolume string) string {
	fromVolumeHash := generateHash(fromVolume)
	return getVolumeJsonFolder(toVolume) + fromVolumeHash + ".json"
}

func getVolumeJsonFolder(toVolume string) string {
	return toVolume + "/.volumes/"
}