package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type ScanData struct {
	ImageDb        string
	RecordedImages map[string]bool
	NewImages      []string
}

var scanData = initialize()
var imageMatcher = imageTypeRegex()

// Create our container that will hold our image processing information
func initialize() *ScanData {
	data := new(ScanData)
	data.ImageDb = "images.txt"
	data.RecordedImages = make(map[string]bool)
	data.NewImages = make([]string, 0)
	return data
}

// Scans the directory for new images
func ScanImages(path string) []string {
	fmt.Printf("Scanning %s for new image files...\n", path)
	readImageDb()
	err := filepath.Walk(path, scan)
	if err != nil {
		fmt.Printf("Errors occured while scanning for new images: %v\n", err)
	}
	writeImageDb()
	return scanData.NewImages
}

// If path is a new image then save it for processing
func scan(path string, fileInfo os.FileInfo, err error) error {
	if isNewImage(path, fileInfo) {
		scanData.NewImages = append(scanData.NewImages, path)
	}
	return nil
}

// Determines if the file is an image and if we have seen it before
func isNewImage(path string, fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir() &&
		imageMatcher.MatchString(filepath.Ext(path)) &&
		!seenImage(path)
}

// Currently we only scan for the following image types: jpeg, jpg or png
func imageTypeRegex() *regexp.Regexp {
	regex, _ := regexp.Compile("(?i).(jpe?g|png)")
	return regex
}

// Determines if we've seen the image before or not
func seenImage(image string) bool {
	if _, seen := scanData.RecordedImages[image]; seen {
		return true
	}
	return false
}

// We keep track of the images we've seen in a flat file database; we load it into memory for lookup
func readImageDb() {
	file, err := os.OpenFile(scanData.ImageDb, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanData.RecordedImages[scanner.Text()] = true
	}
	fmt.Printf("We have previously seen %d images\n", len(scanData.RecordedImages))
}

// Record the new files so we don't process them again
func writeImageDb() {
	numNewImages := len(scanData.NewImages)
	if numNewImages > 0 {
		file, err := os.OpenFile(scanData.ImageDb, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		for _, path := range scanData.NewImages {
			file.WriteString(path + "\n")
		}
	}
}
