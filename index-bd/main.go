package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"emails.com/indexbd/functions"
	"github.com/akamensky/argparse"
)

// función que sube a ZincSearch los datos de las rutas leinas en un archivo
func uploadPararell(wg *sync.WaitGroup, path string, databse string) {
	fmt.Println("reading file: ", path)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		functions.UploadFile(scanner.Text(), databse)
	}
	wg.Done()
}

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	var input *string = parser.String("i", "input", &argparse.Options{Required: true, Help: "input dir"})
	var database *string = parser.String("d", "database", &argparse.Options{Required: true, Help: "database"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string

	functions.MakeFiles(*input)

	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go uploadPararell(&waitGroup, "./files1.txt", *database)
	go uploadPararell(&waitGroup, "./files2.txt", *database)
	go uploadPararell(&waitGroup, "./files3.txt", *database)
	go uploadPararell(&waitGroup, "./files4.txt", *database)
	waitGroup.Wait()

}
