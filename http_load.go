package gwf

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	vegeta "github.com/tsenart/vegeta/lib"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	Body   string        `json:"body"`
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
	var body []byte

	rate := vegeta.Rate{Freq: r.Rate, Per: time.Second}
	duration := 5 * time.Second
	targetUrl := fmt.Sprintf("http://%s%s", url, r.Url)
	fmt.Printf("Testing: %s - %s\n", r.Method, targetUrl)

	if r.Body != "" {
		body = getBody(r.Body)
	}

	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: r.Method,
		URL:    targetUrl,
		Body:   body,
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
	printMetrics(&metrics, &r)
}

func printMetrics(m *vegeta.Metrics, r *LoadRoute) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TYPE", "RESULT"})
	table.Append([]string{"99th percentile", fmt.Sprintf("%s", m.Latencies.P99)})
	table.Append([]string{"Total request", fmt.Sprintf("%d", m.Requests)})
	table.Append([]string{"Duration", fmt.Sprintf("%s", m.Duration)})
	table.Append([]string{"Rate", fmt.Sprintf("%ds", r.Rate)})

	for s, _ := range m.StatusCodes {
		table.Append([]string{fmt.Sprintf("Status code %s", s), strconv.Itoa(m.StatusCodes[s])})
	}

	table.Render()
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

// Read body from .json
func getBody(path string) []byte {
	content, err := ioutil.ReadFile(GetDynamicPath(path))
	if err != nil {
		ProcessError(err)
	}

	return content
}
