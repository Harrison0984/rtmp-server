package main

import (
	"../rtmp"
	"flag"
	"log"
	"runtime"
)

var listen string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	flag.StringVar(&listen, "l", ":1935", "-l=:1935")
}
func main() {
	flag.Parse()
	err := rtmp.ListenAndServe(listen)
	if err != nil {
		panic(err)
	}
	log.Println("rtmp server listen at " + listen)
	select {}
}
