package gh

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

// Fetcher defines the interface for fetching raw content from a URL.
type Fetcher interface {
	Fetch(url string) (string, error)
}

// CURLFetcher implements a way to fetch content using the system's curl command.
type CURLFetcher struct{}

func (f *CURLFetcher) Fetch(url string) (string, error) {
	cmd := exec.Command("curl", "-sL", "-A", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36", url)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("curl error: %v, stderr: %s", err, stderr.String())
	}
	return out.String(), nil
}

// HttpFetcher implements a way to fetch content using Go's net/http.
type HttpFetcher struct{}

func (f *HttpFetcher) Fetch(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// MultiFetcher tries multiple fetchers until one succeeds.
type MultiFetcher struct {
	fetchers []Fetcher
}

func (f *MultiFetcher) Fetch(url string) (string, error) {
	var lastErr error
	for _, fetcher := range f.fetchers {
		content, err := fetcher.Fetch(url)
		if err == nil {
			return content, nil
		}
		lastErr = err
	}
	return "", fmt.Errorf("all fetchers failed: %v", lastErr)
}
