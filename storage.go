package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
	"log"
)

type File struct {
	path string
	isDir bool
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

func removeFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File deleted")
}

func removeFiles(fromVolume string, toVolume string, filesToRemove []string) {
	for _, pathToRemove := range filesToRemove {
		removeFile(pathToRemove)
		removeFromFileCache(pathToRemove, fromVolume, toVolume)

		dirPath := filepath.Dir(pathToRemove)
		filesList := getFilesList(dirPath)
		if (len(filesList) == 0) {
			os.RemoveAll(dirPath)
		}
	}
}

func copyFiles(fromVolume string, toVolume string, filesToCopy []string) {
	volumeName := getVolumeName(fromVolume)
	for _, sourcePath := range filesToCopy {
		pathToFile := strings.Replace(sourcePath, fromVolume, "", -1)
		destinationPath := toVolume + "/" + volumeName + pathToFile
		fmt.Println("Copied " + sourcePath)
		copy(sourcePath, destinationPath)

		sourceFileCreatedAt := getFileCreatedAt(sourcePath)
		destinationFileCreatedAt := getFileCreatedAt(destinationPath)

		fromVolumeHash := generateHash(sourcePath + sourceFileCreatedAt)
		toVolumeHash := generateHash(destinationPath + destinationFileCreatedAt)

		addToFileCache(destinationPath, fromVolume, toVolume, fromVolumeHash, toVolumeHash)
	}
}

func getFileCreatedAt(path string) string {
	destinationFile, err := os.Stat(path)
	if (err != nil) {
		fmt.Println(err)
	}
	var modTime int64 = destinationFile.ModTime().UnixNano()
	return strconv.FormatInt(modTime, 10)
}

func getFilesListRecursive(path string) Files {
	files, err := ioutil.ReadDir(path)
	if (err != nil) {
		log.Fatal(err)
	}

	var filesList Files
	for _, file := range files {
		var modTime int64 = file.ModTime().UnixNano()
		fileStruct := File{path + "/" + file.Name(), file.IsDir(), strconv.FormatInt(modTime, 10)}
		filesList = append(filesList, fileStruct)
	}

	for _, file := range filesList {
		if (file.isDir) {
			for _, fileItem := range getFilesListRecursive(file.path) {
				filesList = append(filesList, fileItem)
			}
		}
	}

	return filesList
}

// TODO:: мб как то объеденить с recursive версией
func getFilesList(path string) Files {
	files, err := ioutil.ReadDir(path)
	if (err != nil) {
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
