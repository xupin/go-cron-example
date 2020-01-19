package main

import "go-cron-example/console"

func main() {

	quitChan := make(chan bool, 1)

	// console
	go func() {
		console.Default()
	}()

	<-quitChan
}
