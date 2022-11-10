package functions

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

// escribe un archivo de texto todas las rutas encontradas en una carpeta
func readfiles(path string) {
	f, err := os.Create("files.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		//Escribe ruta si esta no es un directorio
		if !info.IsDir() {
			fmt.Println(path)
			_, err = f.WriteString(path + "\n")

			if err != nil {
				log.Fatal(err)
			}

		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

//cuenta lineas escritas que tiene un archivo
func lineCounter(r io.Reader) (int, error) {

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

//copia un intervalo de lineas de un archivo a otro
func copyLines(fname string, from, to int, newFile string) {
	new_file, err := os.Create(newFile)
	if err != nil {
		log.Fatal(err)
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
		_, err = new_file.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

// obtiene intarvalos en los que divide la cantidad total de archivos para utilizar los
// 4 nucleos que tiene mi computador
func getSlices(nlines int) []int {
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

// crea los archivos con las rutas a ser procesadas por cada uno de los nucleos
func MakeFiles(path string) {
	readfiles(path)
	file, err := os.Open("files.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	lines, err2 := lineCounter(file)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("lineas: %i", lines)
	slices := getSlices(lines)
	fmt.Println(slices)
	var filenames [4]string
	filenames[0] = "files1.txt"
	filenames[1] = "files2.txt"
	filenames[2] = "files3.txt"
	filenames[3] = "files4.txt"
	for i := 0; i < 4; i++ {
		fmt.Println("slice: [", slices[i], " : ", slices[i+1]-1, "], filename: ", filenames[i])
		copyLines("files.txt", slices[i], slices[i+1]-1, filenames[i])
	}

}
