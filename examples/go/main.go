package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("EKYCPRO_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Set the EKYCPRO_API_KEY environment variable")
		os.Exit(1)
	}

	req, _ := http.NewRequest("POST", "https://api.ekycpro.com/v1/check", bytes.NewBufferString(`{"service_type":"netflix","identifier":"alice@example.com"}`))
	req.Header.Set("X-API-Key", key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "request error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, "request failed: %d %s\n", resp.StatusCode, body)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
