package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	watchName = flag.String("name", ".", "")
	verbose   = flag.Bool("verbose", false, "")
)

func main() {
	flag.Parse()
	command := readString(os.Stdin)
	logPrintf("command %v", command)
	go proc(watch(*watchName), *watchName, command)
	<-(make(chan struct{}))
}

func readString(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("unable to read stdin: %v", err)
	}
	return string(b)
}
