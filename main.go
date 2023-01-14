package main

import (
	"fmt"
)

func init() {
	fromVolume := "/Users/macos/Documents/work/projects/go/kstorage/dir1"
	toVolume := "/Users/macos/Documents/work/projects/go/kstorage/dir2"
	loadToMemory(fromVolume, toVolume)
}

func main() {
	// TODO:: получаем приатаченые девайсы с входа и выхода (from / to)
	// attachedDevices := getAttachedDevices()
	// ??

	fromVolume := "/Users/macos/Documents/work/projects/go/kstorage/dir1"
	toVolume := "/Users/macos/Documents/work/projects/go/kstorage/dir2"
	filesToCopy, filesToRemove := scanDifference(fromVolume, toVolume)

	if (len(filesToCopy) == 0 || len(filesToRemove) == 0) {
		fmt.Println("There are no files to copy or remove")
	}
	fmt.Println(len(filesToRemove), filesToRemove)

	copyFiles(fromVolume, toVolume, filesToCopy)
	removeFiles(fromVolume, toVolume, filesToRemove)
}

// func initFileHashes() {
// 	toVolume := "/Users/macos/Documents/work/projects/go/kstorage/dir2"
// 	filesList := getFilesList(toVolume)

// 	fmt.Println("test", getFilesList("/Users/macos/Documents/work/projects/go/kstorage/dir1"))

// 	for _, file := range filesList {
// 		filePath := strings.Replace(file.path, toVolume, "", -1)
// 		addToMemoryCache(toVolume, filePath)
// 	}
// }
