package gwf

import (
	"encoding/json"
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"io/ioutil"
	"net/http"
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
	Header string        `json:"header"`
	Time   time.Duration `json:"duration"`
	Rate   int           `json:"rate"`
}

// Command registration
func (c *HttpLoad) Register() {
	c.Signature = "http:load target.json" // Change command signature
	c.Description = "Execute http load"   // Change command description
}

// Command business logic
func (c *HttpLoad) Run(conf *Conf) {
	var routes FileStruct
	var serverName string
	readJsonFile(c.Args, &routes)

	if conf.Server.Name == "" {
		serverName = "localhost"
	} else {
		serverName = conf.Server.Name
	}

	for _, r := range routes.Routes {
		attack(r, fmt.Sprintf("%s:%d", serverName, conf.Server.Port))
	}
}

// Execute Vegeta attack
func attack(r LoadRoute, url string) {
	rate := vegeta.Rate{Freq: r.Rate, Per: time.Second}
	duration := 5 * time.Second
	targetUrl := fmt.Sprintf("http://%s%s", url, r.Url)
	fmt.Printf("Testing: %s\n", targetUrl)

	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: r.Method,
		URL:    targetUrl,
		Body:   r.Body,
		Header: http.Header{
			"Content-Type": {
				r.Header,
			},
		},
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
func readJsonFile(path string, str *FileStruct) {
	filePath := GetDynamicPath(path)
	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(jsonFile, &str); err != nil {
		ProcessError(err)
	}
}
