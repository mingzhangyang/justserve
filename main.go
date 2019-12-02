package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var dir string
	var port int
	var err error

	switch len(os.Args) {
	case 4:
		if os.Args[1] == "-p" {
			dir = os.Args[3]
			port, err = strconv.Atoi(os.Args[2])
		}
		if os.Args[2] == "-p" {
			dir = os.Args[1]
			port, err = strconv.Atoi(os.Args[3])
		}
		if err != nil {
			log.Fatal("invalid port number")
		}
		if port < 0 || port > 65535 {
			log.Fatal("invalid port number")
		}
	case 2:
		dir = os.Args[1]
		port = 9102
	default:
		log.Println("invalid arguments.")
		printUsage()
		os.Exit(1)
	}

	info, err := os.Stat(dir)
	if err != nil {
		log.Fatal("invalid directory to serve")
	}

	if !info.IsDir() {
		log.Fatal(dir + " is not a directory")
	}

	fmt.Printf("Serving the files in %s\n", http.Dir(dir))
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Printf("Listening on port: %d", port)
	log.Printf("Please visit http://localhost:%d", port)
	log.Println("Ctrl+C to quit.")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func printUsage() {
	fmt.Print("Usage: \n    ./justserve path/to/the/directory/to/be/served\n")
	fmt.Print("    ./justserve -p 8888 path/to/the/directory/to/be/served\n")
}
