package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	tmpDir := flag.String("tmpdir", "/tmp", "directory to temporarily write downloaded files to")
	method := flag.String("method", "get", "HTTP method")
	urlstr := flag.String("url", "", "the URL to act upon")
	dest := flag.String("dest", "", "the destination for the retrieved file (when method=get)")
	overwrite := flag.Bool("overwrite", false, "when true, will overwrite the file specified by -dest")
	flag.Parse()

	// Check to see if the temporary directory exists, and is a directory.
	if pathExits(*tmpDir) {
		if !isDir(*tmpDir) {
			log.Fatalf("%q exists, but is not a directory", *tmpDir)
		}
	} else {
		log.Fatalf("%q does not exist", *tmpDir)
	}

	// Check to see if the path set by -dest= exists, and if it is,
	if pathExists(*dest) && !*overwrite {
		log.Fatalln("path exists, but not allowed to overwrite")
	}

}

func pathExists(pth string) bool {
	_, err := os.Stat(pth)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
