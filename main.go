package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// the default config file location that ubuntu installs for google chrome
const defaultFile = "/usr/share/applications/google-chrome.desktop"

func main() {
	// allow the config file to be applied using flags
	file := flag.String("file", defaultFile, "The file location of your google-chrome.desktop config file")

	// read the config file, error in fail
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("error opening file '%s': %v", *file, err)
	}

	// do some string replacement to add the force-dark-mode flag to google chrome
	str := string(data)
	str = strings.Replace(str, "Exec=/usr/bin/google-chrome-stable %U\n", "Exec=/usr/bin/google-chrome-stable %U --force-dark-mode\n", 1)
	str = strings.Replace(str, "Exec=/usr/bin/google-chrome-stable\n", "Exec=/usr/bin/google-chrome-stable --force-dark-mode\n", 1)

	// write old file first, if we can't do this then dont write the new stuff
	// this gives us a backup of our original config file (just to be safe!)
	err = ioutil.WriteFile(*file+".old", data, 0644)
	if err != nil {
		log.Fatalf("failed writing backup file '%s': %v", *file+".old", err)
	} else {
		// backup config has been written, lets go ahead and write the new config with dark mode flag enabled
		err := ioutil.WriteFile(*file, []byte(str), 0644)
		if err != nil {
			log.Fatalf("failed writing file '%s': %v", *file, err)
		}
	}

	// all done
	fmt.Println(("Done. Please restart your system for changes to take effect."))
}
