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
		fmt.Printf("Found %d new images\n", len(newImages))
		for _, image := range newImages {
			fmt.Printf("%s\n", image)
		}
		if doUpload(newImages) {
			fmt.Printf("Attempting to upload all %d images...\n", len(newImages))
		}
	} else {
		fmt.Printf("Did not find any new images to upload\n")
	}
}

// Ask the user if they want to upload all found new images
func doUpload(newImages []string) bool {
	fmt.Printf("Would you like to upload all %d images to your Flickr account? Y or N?\n", len(newImages))
	var answer string
	fmt.Scanf("%s", &answer)
	return strings.ToLower(strings.TrimSpace(answer)) == "y"
}
