package gwf

import (
	"encoding/json"
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"net/http"
	"os"
	"time"
)

type HttpLoad struct {
	Signature   string
	Description string
	Args        string
}

type FileStruct struct {
	Routes []LoadRoute `json:"routes"`
}

type LoadRoute struct {
	Method string        `json:"method"`
	Url    string        `json:"url"`
	Body   []byte        `json:"body"`
	Header http.Header   `json:"header"`
	Time   time.Duration `json:"duration"`
	Rate   int           `json:"rate"`
}

// Command registration
func (c *HttpLoad) Register() {
	c.Signature = "http:load <target.json>" // Change command signature
	c.Description = "Execute http load"     // Change command description
}

// Command business logic
func (c *HttpLoad) Run(conf *Conf) {
	var routes FileStruct
	readJsonFile(c.Args, &routes)

	for _, r := range routes.Routes {
		attack(r, fmt.Sprintf("%s:%b", conf.Server.Name, conf.Server.Port))
	}
}

// Execute Vegeta attack
func attack(r LoadRoute, url string) {
	rate := vegeta.Rate{Freq: r.Rate, Per: time.Second}
	duration := r.Time

	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: r.Method,
		URL:    fmt.Sprintf("%s/%s", url, r.Url),
		Body:   r.Body,
		Header: r.Header,
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(target, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

// Read JSON file content
func readJsonFile(path string, str *FileStruct) *os.File {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal([]byte(path), &str); err != nil {
		ProcessError(err)
	}

	return jsonFile
}
