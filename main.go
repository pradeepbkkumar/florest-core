package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var projectPath = flag.String("projectPath", "", "path to project directory e.g /home/user/myapp")

const (
	florest_core  = "github.com/jabong/florest-core"
	newapp_script = "scripts/newapp.sh"
)

func main() {
	flag.Parse()
	if *projectPath == "" {
		flag.PrintDefaults()
		return
	}
	var err error
	// check if project path exists
	if _, err = os.Stat(*projectPath); os.IsNotExist(err) {
		log.Println("projectPath:" + *projectPath + " not found")
		return
	}

	// check if gopath is set
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Println("Please, set $GOPATH environment variable\n")
		return
	}

	//Support gopaths with multiple directories
	dirs := strings.Split(gopath, ":")
	var fullPath string
	found := false
	for _, d := range dirs {
		fullPath = path.Join(d, "src", florest_core)
		if _, err = os.Stat(fullPath); err == nil {
			found = true
			break
		}
	}
	// return if florest was not found
	if found == false {
		log.Fatal(florest_core + " not found")
	}
	// florest-core is present, make the new app now
	out, err := exec.Command("/bin/sh", path.Join(fullPath, newapp_script), fullPath, *projectPath).Output()
	if err != nil {
		log.Println("error in creating new app:" + err.Error())
	} else {
		log.Println(string(out))
	}
}
