package main

import (
	// "dkrylov/kstorage-processing/src/config"
	"dkrylov/kstorage-processing/src/cache"
	"dkrylov/kstorage-processing/src/storage"
	"strings"
	"flag"
	"fmt"
)

func getNoIndex() []string {
	noIndex := flag.String("noIndex", "", "This paths will be not processed")
	flag.Parse()
	return strings.Split(*noIndex, ",")
}

func init() {
	// TODO::Config loader (at least for volumesCacheDirName)
	// conf := config.Load2("/Users/dmitriykrylov/Documents/work/projects/go/kstorage-processing/config.yaml")
	// fmt.Println(conf)

	// config.Load("/Users/dmitriykrylov/Documents/work/projects/go/kstorage-processing/config.yaml")
}

func main() {
	fromVolume := flag.String("fromVolume", "", "Path to source volume")
	toVolume := flag.String("toVolume", "", "Path to destination volume")
	disableDeletion := flag.Bool("disableDeletion", false, "Disable deletion of files that has been removed from the source storage")
	noIndex := getNoIndex()
	flag.Parse()

	if (len(*fromVolume) == 0) {
		fmt.Println("Parameter `fromVolume` is not set")
		return
	}

	if (len(*toVolume) == 0) {
		fmt.Println("Parameter `toVolume` is not set")
		return
	}

	cache.LoadToMemory(*fromVolume, *toVolume)

	filesToCopy, filesToRemove := storage.ScanDifference(*fromVolume, *toVolume, noIndex)

	if len(filesToCopy) == 0 {
		fmt.Println("There are no files to copy")
		return
	}
	storage.CopyFiles(*fromVolume, *toVolume, filesToCopy)

	if len(filesToRemove) == 0 {
		fmt.Println("There are no files to remove")
		return
	}

	if (!*disableDeletion) {
		storage.RemoveFiles(*fromVolume, *toVolume, filesToRemove)
	}
}
