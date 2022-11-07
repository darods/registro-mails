package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	took      string
	timed_out string
}

func simpleMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World!"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode("You got me")
	palabra := "You got me"
	w.Write([]byte(palabra))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("POSTZZZZ")
}

func nameExample(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	lastname := r.URL.Query().Get("lastname")
	w.Write([]byte("hi " + lastname + " " + name))
}

func getZincBasic(w http.ResponseWriter, r *http.Request) {
	term := "USA"
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `",
            "start_time": "2022-10-01T14:28:31.894Z",
            "end_time": "2022-10-24T15:28:31.894Z"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/olympics/_search", strings.NewReader(query))
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func getZincSearch(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `",
            "start_time": "2022-10-01T14:28:31.894Z",
            "end_time": "2022-10-24T15:28:31.894Z"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
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
	w.Header().Set("Content-Type", "application/json")
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	/*
		for _, rec := range result.Hits.Hits {
			fmt.Println(rec.Index)
			json.NewEncoder(w).Encode(rec)
		}*/
	json.NewEncoder(w).Encode(result.Hits.Hits)
	fmt.Println(string(body))

}

//https://dev.to/billylkc/parse-json-api-response-in-go-10ng
//https://mholt.github.io/json-to-go/
type Response struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
				ContentType             string `json:"Content-Type"`
				Date                    string `json:"Date"`
				From                    string `json:"From"`
				Message                 string `json:"Message"`
				MessageID               string `json:"Message-ID"`
				MimeVersion             string `json:"Mime-Version"`
				Subject                 string `json:"Subject"`
				To                      string `json:"To"`
				XFileName               string `json:"X-FileName"`
				XFolder                 string `json:"X-Folder"`
				XFrom                   string `json:"X-From"`
				XOrigin                 string `json:"X-Origin"`
				XTo                     string `json:"X-To"`
				XBcc                    string `json:"X-bcc"`
				XCc                     string `json:"X-cc"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
