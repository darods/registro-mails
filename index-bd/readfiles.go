package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func readfiles(path string) {
	f, err1 := os.Create("files.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		//fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		// Si no es un directorio
		if !info.IsDir() {
			fmt.Println(path)
			_, err2 := f.WriteString(path + "\n")

			if err2 != nil {
				log.Fatal(err2)
			}

		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func LineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func CopyLines(fname string, from, to int, file string) {
	n_f, err1 := os.Create(file)
	if err1 != nil {
		log.Fatal(err1)
	}
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n++
		if n < from {
			continue
		}
		if n > to {
			break
		}
		_, err2 := n_f.WriteString(scanner.Text() + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}

}

func getSeparation(nlines int) []int {
	numCpu := runtime.NumCPU()
	var slices [5]int
	slices[0] = 0
	division := nlines / (numCpu + 1)
	for i := 1; i < numCpu+1; i++ {
		if i+1 < numCpu {
			slices[i] = division * (i + 1)
		} else {
			slices[i] = (division * (i + 1)) + (nlines % (numCpu + 1))
		}

	}
	return slices[:]
}

func main() {
	fmt.Printf("analizando carpeta")
	readfiles("./enron_mail_20110402/")
	file, err := os.Open("files.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	lines, err2 := LineCounter(file)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("lineas: %i", lines)
	slices := getSeparation(lines)
	fmt.Println(slices)
	var filenames [4]string
	filenames[0] = "files1.txt"
	filenames[1] = "files2.txt"
	filenames[2] = "files3.txt"
	filenames[3] = "files4.txt"
	for i := 0; i < 4; i++ {
		fmt.Println("slice: [", slices[i], " : ", slices[i+1]-1, "], filename: ", filenames[i])
		CopyLines("files.txt", slices[i], slices[i+1]-1, filenames[i])
	}

}
