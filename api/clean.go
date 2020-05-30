package api

import (
	"log"
	"os"
	"strings"
)

var dir string = "/tmp"

// Clean delete ClI downloaded files
func Clean() {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Failed open dir %v: %v", dir, err)
	}
	defer d.Close()
	names, err := d.Readdir(-1)
	if err != nil {
		log.Fatalf("Failed to list dir %v: %v", dir, err)
	}

	for _, n := range names {
		if strings.HasPrefix(n.Name(), "GetGit") {
			absName := dir + "/" + n.Name()
			err := os.Remove(absName)
			if err != nil {
				log.Fatalf("Failed to delete %v: %v", absName, err)
			}
		}
	}
}
