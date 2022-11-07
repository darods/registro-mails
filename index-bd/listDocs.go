package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func insertData(filename string) {
	app := "curl"
	arg0 := "http://localhost:4080/api/_bulk"
	arg1 := "-i"
	arg2 := "-u"
	arg3 := "admin:Complexpass#123"
	arg4 := "--data-binary"
	arg5 := "@" + filename

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stdout))
	fmt.Println(filename + " insertado con exito")

	deleteFile(filename)

}

func writeNDJson(filename string, jsonStr []byte) {
	dataInfo := "{ \"index\" : { \"_index\" : \"emails\" } } "
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(dataInfo + "\n"); err != nil {
		log.Println(err)
	}
	if _, err := f.Write(jsonStr); err != nil {
		log.Println(err)
	}
}

func deleteFile(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	err := filepath.Walk("./enron_mail_20110402/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		//fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		// Si no es un directorio
		if !info.IsDir() {
			file, err := os.Open(path)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			var wordlines []string
			for scanner.Scan() {
				wordlines = append(wordlines, scanner.Text())
			}

			dict := map[string]string{
				"Message-ID":                "",
				"Date":                      "",
				"From":                      "",
				"To":                        "",
				"Subject":                   "",
				"Mime-Version":              "",
				"Content-Type":              "",
				"Content-Transfer-Encoding": "",
				"X-From":                    "",
				"X-To":                      "",
				"X-cc":                      "",
				"X-bcc":                     "",
				"X-Folder":                  "",
				"X-Origin":                  "",
				"X-FileName":                "",
				"Message":                   "",
			}
			compareString := []string{"Message-ID", "Date", "From", "To", "Subject", "Mime-Version", "Content-Type", "Content-Transfer-Encoding", "X-From", "X-To", "X-cc", "X-bcc", "X-Folder", "X-Origin", "X-FileName"}

			aux := ""
			for _, words := range wordlines {
				msg := compareString[0]
				res := strings.Index(words, msg)

				if res == 0 {
					words = strings.Replace(words, msg+": ", "", 1)
					dict[msg] += words
					aux = msg
					if len(compareString) != 1 {
						compareString = compareString[1:]
					} else {
						compareString[0] = "Message"
						aux = "Message"
					}

				} else {
					dict[aux] += words
				}
			}

			jsonStr, err := json.Marshal(dict)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			writeNDJson(path+".ndjson", jsonStr)
			insertData(path + ".ndjson")

		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
