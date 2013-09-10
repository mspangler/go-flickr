package main

import (
	"flag"
	"fmt"
	"github.com/mspangler/go-flickr/io"
	"os"
	"runtime/pprof"
)

const (
	Image_Directory = "/Users/mark/Pictures/"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

// Start the application
func main() {

	newImages := io.ScanImages(Image_Directory)
	numNewImages := len(newImages)
	if numNewImages > 0 {
		fmt.Printf("Will try to upload %d new images\n", numNewImages)
	} else {
		fmt.Printf("Did not find any new images to upload\n")
	}
	// TODO: ask the user to upload or not; don't want to upload if they already have; maybe only care about new ones from this point on

	flag.Parse()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
