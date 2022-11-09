package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

// Task that will be done over time
func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1.")
	wg.Done()
}
func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2.")
	wg.Done()
}
func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3.")
	wg.Done()
}

// Task done in parallel
func listenForever() {
	for {
		printTime("Listening...")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	waitGroup.Add(4)

	// Give some time for listenForever to start
	time.Sleep(time.Nanosecond * 10)

	// Let's start writing the mails
	go writeMail1(&waitGroup)
	go writeMail2(&waitGroup)
	go writeMail3(&waitGroup)
	go writeMail3(&waitGroup)

	waitGroup.Wait()
}
