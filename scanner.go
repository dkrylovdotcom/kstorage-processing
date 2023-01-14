package main

import (
	"fmt"
	// "strings"
	"github.com/deepakjois/gousbdrivedetector"
)

func getAttachedDevices() []string {
	if drives, err := usbdrivedetector.Detect(); err == nil {
		fmt.Printf("%d USB Devices Found\n", len(drives))
		return drives
	} else {
		fmt.Println(err)
		return nil
	}
}

func scanDifference(fromVolume string, toVolume string) ([]string, []string) {
	filesList := getFilesListRecursive(fromVolume)
	filesToCopy := getFilesToCopy(filesList, fromVolume, toVolume)
	filesToRemove := getFilesToRemove(filesList, fromVolume, toVolume)
	return filesToCopy, filesToRemove
}

func getFilesToRemove(filesList []File, fromVolume string, toVolume string) []string {
	var filesToRemove []string
	if (!isVolumeJsonExists(fromVolume, toVolume)) {
		saveFile(fromVolume, toVolume, []byte ("{}"))
	}
	volumeJson := getVolumeJson(fromVolume, toVolume)
	for _, cacheItem := range volumeJson {
		isExists := false
		for _, file := range filesList {
			if (file.isDir) {
				continue
			}

			fileHash := generateHash(file.path + file.updatedAt)
			if (fileHash == cacheItem.FromVolumeHash) {
				isExists = true
			}
		}
		if (!isExists) {
			pathToRemove := cacheItem.Path
			filesToRemove = append(filesToRemove, pathToRemove)
		}
	}
	return filesToRemove
}

func getFilesToCopy(filesList []File, fromVolume string, toVolume string) []string {
	var filesToCopy []string
	for _, file := range filesList {
		if (file.isDir) {
			continue
		}

		fileHash := generateHash(file.path + file.updatedAt)
		fileExistsAndNotChanged := findInMemory(toVolume, fileHash)
		if (!fileExistsAndNotChanged) {
			filesToCopy = append(filesToCopy, file.path)
		}
	}
	return filesToCopy
}
