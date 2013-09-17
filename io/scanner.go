package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type ScanData struct {
	Database        string
	Regex           *regexp.Regexp
	ImageCollection map[string]bool
	NewImages       []string
}

var scanData = newScanData()

// Create our container that will hold our image processing information
func newScanData() *ScanData {
	regex, _ := regexp.Compile("(?i).(jpe?g|png)")
	return &ScanData{
		Database:        "db",
		Regex:           regex,
		ImageCollection: make(map[string]bool),
		NewImages:       make([]string, 0),
	}
}

// Scans the directory for new images
func ScanImages(path string) []string {
	fmt.Printf("Scanning %s for new image files...\n", path)
	readDatabase()
	err := filepath.Walk(path, scan)
	if err != nil {
		fmt.Printf("Errors occured while scanning for new images: %v\n", err)
	}
	writeDatabase()
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
		scanData.Regex.MatchString(filepath.Ext(path)) &&
		!scanData.ImageCollection[path]
}

// We keep track of the images we've seen in a flat file database; we load it into memory for lookup
func readDatabase() {
	file, err := os.OpenFile(scanData.Database, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanData.ImageCollection[scanner.Text()] = true
	}
	fmt.Printf("We have previously seen %d images\n", len(scanData.ImageCollection))
}

// Record the new files so we don't process them again
func writeDatabase() {
	numNewImages := len(scanData.NewImages)
	fmt.Printf("Found %d new images\n", numNewImages)
	if numNewImages > 0 {
		file, err := os.OpenFile(scanData.Database, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		for _, path := range scanData.NewImages {
			file.WriteString(path + "\n")
		}
	}
}
