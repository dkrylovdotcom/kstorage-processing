package storage

import (
	"dkrylov/kstorage-processing/src/cache"
)

func ScanDifference(
	fromVolume string,
	toVolume string,
	noIndex []string,
) ([]string, []string) {
	filesList := getFilesListRecursive(fromVolume, noIndex)
	filesToCopy := getFilesToCopy(filesList, fromVolume, toVolume, noIndex)
	filesToRemove := getFilesToRemove(filesList, fromVolume, toVolume, noIndex)
	return filesToCopy, filesToRemove
}

func getFilesToRemove(
	filesList []File,
	fromVolume string,
	toVolume string,
	noIndex []string,
) []string {
	var filesToRemove []string
	fileCache := cache.GetFileCache(fromVolume, toVolume)
	for _, cacheItem := range fileCache {
		isExists := false
		for _, file := range filesList {
			if file.isDir {
				continue
			}
			IsPathPresent := IsPathPresent(noIndex, file.path)
			if (IsPathPresent) {
				continue
			}

			fileHash := cache.GenerateHash(file.path + file.updatedAt)
			if fileHash == cacheItem.FromVolumeHash {
				isExists = true
			}
		}
		if !isExists {
			pathToRemove := cacheItem.Path
			filesToRemove = append(filesToRemove, pathToRemove)
		}
	}
	return filesToRemove
}

func getFilesToCopy(
	filesList []File,
	fromVolume string,
	toVolume string,
	noIndex []string,
) []string {
	var filesToCopy []string
	for _, file := range filesList {
		if file.isDir {
			continue
		}
		IsPathPresent := IsPathPresent(noIndex, file.path)
		if (IsPathPresent) {
			continue
		}

		fileHash := cache.GenerateHash(file.path + file.updatedAt)
		fileExistsAndNotChanged := cache.FindInMemory(toVolume, fileHash)
		if !fileExistsAndNotChanged {
			filesToCopy = append(filesToCopy, file.path)
		}
	}
	return filesToCopy
}
