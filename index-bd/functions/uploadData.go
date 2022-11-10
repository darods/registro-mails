package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// sube directo informacion json a zincsearch
func writeDirect(jsonStr []byte, database string) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+database+"/_doc", strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func UploadFile(path string, database string) {

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
	writeDirect(jsonStr, database)

}

/*
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
	dataInfo := "{ \"index\" : { \"_index\" : \"emails_2\" } } "
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
*/
