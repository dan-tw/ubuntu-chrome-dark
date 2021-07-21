package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const file = "/usr/share/applications/google-chrome.desktop"

func main() {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error opening file '%s': %v", file, err)
	}

	str := string(data)
	str = strings.Replace(str, "Exec=/usr/bin/google-chrome-stable %U\n", "Exec=/usr/bin/google-chrome-stable %U --force-dark-mode\n", 1)
	str = strings.Replace(str, "Exec=/usr/bin/google-chrome-stable\n", "Exec=/usr/bin/google-chrome-stable --force-dark-mode\n", 1)

	// write old file first, if we can't do this then dont write the new stuff
	err = ioutil.WriteFile(file+".old", data, 0644)
	if err != nil {
		log.Fatalf("failed writing backup file '%s': %v", file+".old", err)
	} else {
		err := ioutil.WriteFile(file, []byte(str), 0644)
		if err != nil {
			log.Fatalf("failed writing file '%s': %v", file, err)
		}
	}

	fmt.Println(("Done. Please restart your system for changes to take effect."))
}
