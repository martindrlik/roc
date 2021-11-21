package main

import "log"

func logPrintf(format string, v ...interface{}) {
	if *verbose {
		log.Printf(format, v...)
	}
}
