package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	method := flag.String("method", "get", "HTTP method")
	urlstr := flag.String("url", "", "the URL to act upon")
	dest := flag.String("dest", "", "the destination for the retrieved file (when method=get)")
	overwrite := flag.Bool("overwrite", false, "when true, will overwrite the file specified by -dest")
	flag.Parse()

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
