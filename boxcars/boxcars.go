package main

import (
	"flag"
	"fmt"
	"github.com/jmcarbo/boxcars"
	"github.com/jmcarbo/boxcars/json-config"
	"os"
)

var (
	filename string
	port     string
	user_id  int
	group_id int
	secure   bool
)

func main() {
	flag.StringVar(&port, "port", ":8080", "Port to listen")
	flag.BoolVar(&secure, "secure", false, "Enables secure mode to avoid running as sudo.")
	flag.IntVar(&user_id, "uid", 1000, "User id that'll own the system process.")
	flag.IntVar(&group_id, "gid", 1000, "Group id that'll own the system process.")
	flag.Parse()

	filename = flag.Arg(0)

	if filename == "" {
		fmt.Printf("Usage: boxcars config.json\n")
		os.Exit(1)
	}

	go func () {
		config := JSONConfig.NewJSONConfig(filename, boxcars.SetupSites)
		config.EnableAutoReload()
	}()

	if secure {
		go boxcars.Secure(user_id, group_id)
	}

	boxcars.Listen(port)
}
