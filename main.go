package main

import (
	"log"
	"flag"
	"net/http"
	"os"
	"fmt"
	"strconv"
)

var port string

func init() {
	flag.StringVar(&port, "p", "6080", "set the port")
}

func main() {
	flag.Parse()
	// log.Println(port)
	// log.Println(len(os.Args))
	// fmt.Println(os.Args)
	var dir string
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("invalid port argument: not an int\n")
	}
	if p < 0 || p > 65536 {
		log.Fatal("invalid port argument: out of range\n")
	}
	// no custom port provided
	if port == "6080" && len(os.Args) == 2 {
		dir = os.Args[1];
		// fmt.Printf("dir: %s\n", dir)
	}
	// custom port provided
	if len(os.Args) == 4 {
		if os.Args[1] == "-p" {
			dir = os.Args[3]
		} else if os.Args[2] == "-p" {
			dir = os.Args[1]
		}
		// fmt.Printf("dir: %s\n", dir)
	}
	if len(os.Args) != 4 && len(os.Args) != 2 {
		log.Fatal("invalid arguments.\n")
	}
	fmt.Printf("Serving the files in %s\n", http.Dir(dir))
	fs := http.FileServer(http.Dir(dir))
  	http.Handle("/", fs)

	log.Println("Listening on port: " + port)
	log.Println("Please visit http://localhost:"+port)
	log.Println("Ctrl+C to quit.")
  	log.Fatal(http.ListenAndServe(":"+port, nil))
}
