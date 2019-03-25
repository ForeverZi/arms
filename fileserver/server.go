package main

import (
	"net/http"
	"flag"
	"path/filepath"
	"log"
	"strconv"
	"os"
)

var port = flag.Int("p", 17777, "set listen port, default is 17777")
var dir = flag.String("d", ".", "set listen dir, default is pwd")

func main(){
	flag.Parse()
	maxPort := 1 << 16 -1
	if *port > maxPort {
		log.Fatal("invalid prot, max is ", maxPort)
	}
	fi, err := os.Lstat(*dir)
	if err != nil {
		log.Fatal(err)
	}
	if ! fi.IsDir() {
		apath, _ := filepath.Abs(*dir)
		log.Fatal(apath, " not a dir")
	}
	addr := ":"+strconv.FormatInt(int64(*port), 10)
	log.Println("listen port:", *port)
	log.Println("listen dir:", *dir)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(*dir))))
}