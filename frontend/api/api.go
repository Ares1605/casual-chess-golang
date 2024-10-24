package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Colors for terminal output
const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorWhite  = "\033[37m"
)

// prettyJSON formats JSON with indentation
func prettyJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "  ")
	if err != nil {
		return string(data)
	}
	return prettyJSON.String()
}

// startLogGroup starts a new logged group
func startLogGroup() {
	fmt.Printf("\n%s=== HTTP Request/Response Log ===%s\n", colorWhite, colorReset)
}

// endLogGroup ends the logged group
func endLogGroup() {
	fmt.Printf("%s═══════════════════════════════%s\n", colorWhite, colorReset)
}

// logRequest prints request details
func logRequest(req *http.Request, data *url.Values) {
	fmt.Printf("%s=== REQUEST ===%s\n", colorYellow, colorReset)
	fmt.Printf("%sMethod:%s %s\n", colorBlue, colorReset, req.Method)
	fmt.Printf("%sURL:%s %s\n", colorBlue, colorReset, req.URL.String())
	
	// Print Headers
	fmt.Printf("%sHeaders:%s\n", colorBlue, colorReset)
	for key, values := range req.Header {
		fmt.Printf("  %s: %s\n", key, strings.Join(values, ", "))
	}
	
	// Print Query Parameters
	if data != nil && len(*data) > 0 {
		fmt.Printf("%sQuery Parameters:%s\n", colorBlue, colorReset)
		for key, values := range *data {
			fmt.Printf("  %s: %s\n", key, strings.Join(values, ", "))
		}
	}
}

// logResponse prints response details
func logResponse(resp *http.Response, body []byte, duration time.Duration) {
	fmt.Printf("\n%s=== RESPONSE ===%s\n", colorGreen, colorReset)
	fmt.Printf("%sStatus:%s %s\n", colorPurple, colorReset, resp.Status)
	fmt.Printf("%sDuration:%s %v\n", colorPurple, colorReset, duration)
	
	// Print Headers
	fmt.Printf("%sHeaders:%s\n", colorPurple, colorReset)
	for key, values := range resp.Header {
		fmt.Printf("  %s: %s\n", key, strings.Join(values, ", "))
	}
	
	// Print Body
	if len(body) > 0 {
		fmt.Printf("%sBody:%s\n", colorPurple, colorReset)
		if resp.Header.Get("Content-Type") == "application/json" {
			fmt.Println(prettyJSON(body))
		} else {
			fmt.Println(string(body))
		}
	}
}

func Get[T any](endpoint string, data *url.Values, token *string) (*T, error) {
	client := &http.Client{}
	if data == nil {
		data = &url.Values{}
	}
	
	apiUrl := "http://localhost:8080/" + endpoint
	if data != nil && len(*data) > 0 {
		apiUrl += "?" + data.Encode()
	}
	
	req, err := http.NewRequest("GET", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Add("Content-Type", "application/json")
	if token != nil {
		req.Header.Add("Authorization", "Bearer " + *token)
	}
	
	// Start log group and log request
	startLogGroup()
	logRequest(req, data)
	
	// Track request duration
	start := time.Now()
	
	// Send request
	resp, err := client.Do(req)
	if err != nil {
		endLogGroup()
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		endLogGroup()
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	
	// Calculate duration and log response
	duration := time.Since(start)
	logResponse(resp, body, duration)
	endLogGroup()
	
	// Parse response
	var parsed T
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	
	return &parsed, nil
}
