package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func watch(name string) <-chan string {
	logPrintf("creating new watcher")
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("unable to create new watcher: %v", err)
	}

	isWrite := func(op fsnotify.Op) bool {
		return op&fsnotify.Create == fsnotify.Create ||
			op&fsnotify.Write == fsnotify.Write
	}

	ch := make(chan string)
	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				logPrintf("received event: %v", event)
				if isWrite(event.Op) {
					ch <- event.Name
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				logPrintf("received error: %v", err)
			}
		}
	}()

	err = w.Add(name)
	if err != nil {
		log.Fatalf("unable to start watching file or directory %q", err)
	}
	return ch
}
