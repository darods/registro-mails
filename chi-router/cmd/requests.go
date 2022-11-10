package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetZincSearch(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "` + term + `",
            "start_time": "2022-10-01T14:28:31.894Z",
            "end_time": "2022-11-24T15:28:31.894Z"
        },
        "from": ` + from + `,
        "max_results": ` + to + `,
        "_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails3/_search", strings.NewReader(query))
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
	json.NewEncoder(w).Encode(result.Hits.Hits)

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
