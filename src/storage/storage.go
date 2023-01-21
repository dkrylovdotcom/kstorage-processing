package storage

import (
	"dkrylov/kstorage-processing/src/cache"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
	path      string
	isDir     bool
	updatedAt string
}

type Files = []File

func copy(sourceFile string, destinationFile string) {
	input, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := filepath.Base(destinationFile)
	pathToFile := strings.Replace(destinationFile, fileName, "", -1)
	os.MkdirAll(pathToFile, os.ModePerm)

	err = os.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}

func removeFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File deleted", filePath)
}

func RemoveFiles(fromVolume string, toVolume string, filesToRemove []string) {
	for _, pathToRemove := range filesToRemove {
		removeFile(pathToRemove)
		cache.RemoveFromFileCache(pathToRemove, fromVolume, toVolume)

		dirPath := filepath.Dir(pathToRemove)
		filesList := getFilesList(dirPath)
		if len(filesList) == 0 {
			os.RemoveAll(dirPath)
		}
	}
}

func CopyFiles(fromVolume string, toVolume string, filesToCopy []string) {
	volumeName := getVolumeName(fromVolume)
	for _, sourcePath := range filesToCopy {
		pathToFile := strings.Replace(sourcePath, fromVolume, "", -1)
		destinationPath := toVolume + "/" + volumeName + pathToFile
		fmt.Println("File copied" + sourcePath)
		copy(sourcePath, destinationPath)

		sourceFileCreatedAt := getFileCreatedAt(sourcePath)
		destinationFileCreatedAt := getFileCreatedAt(destinationPath)

		fromVolumeHash := cache.GenerateHash(sourcePath + sourceFileCreatedAt)
		toVolumeHash := cache.GenerateHash(destinationPath + destinationFileCreatedAt)

		cache.AddToFileCache(destinationPath, fromVolume, toVolume, fromVolumeHash, toVolumeHash)
	}
}

func IsPathPresent(strings []string, x string) bool {
	for _, n := range strings {
		if n == x {
			return true
		}
	}
	return false
}

func getFileCreatedAt(path string) string {
	destinationFile, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	var modTime int64 = destinationFile.ModTime().UnixNano()
	return strconv.FormatInt(modTime, 10)
}

func getFilesListRecursive(path string, noIndex []string) Files {
	filesList := getFilesList(path)

	for _, file := range filesList {
		IsPathPresent := IsPathPresent(noIndex, file.path)
		if (IsPathPresent) {
			continue
		}
		if (file.isDir) {
			filesInChildDir := getFilesListRecursive(file.path, noIndex)
			filesList = append(filesList, filesInChildDir...)
		}
	}

	return filesList
}

func getFilesList(path string) Files {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var filesList Files
	for _, file := range files {
		var modTime int64 = file.ModTime().UnixNano()
		fileStruct := File{path + "/" + file.Name(), file.IsDir(), strconv.FormatInt(modTime, 10)}
		filesList = append(filesList, fileStruct)
	}

	return filesList
}

func getVolumeName(volume string) string {
	return filepath.Base(volume)
}
