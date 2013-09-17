package main

import (
	"fmt"
	"github.com/mspangler/go-flickr/io"
	"strings"
)

const (
	Image_Directory = "/Users/mark/Pictures/craigslist"
)

// Start the application
func main() {
	newImages := io.ScanImages(Image_Directory)
	numNewImages := len(newImages)
	if numNewImages > 0 {
		for _, image := range newImages {
			// TODO: maybe generate a HTML page to view the new images?
			fmt.Printf("%s\n", image)
		}
		if doUpload(newImages) {
			fmt.Printf("Attempting to upload all %d images...\n", numNewImages)
			// TODO: authenticate with Flickr & upload images
		}
	}
}

// Ask the user if they want to upload all found new images
func doUpload(newImages []string) bool {
	fmt.Printf("Would you like to upload all %d images to your Flickr account? Y or N?\n", len(newImages))
	var answer string
	fmt.Scanf("%s", &answer)
	return strings.ToLower(strings.TrimSpace(answer)) == "y"
}
