package request

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Request represents a single request from request.txt
type Request struct {
	Prefix string 
	From   string 
	To     string 
	Raw    string // the original line
}

// ParseRequestLine parses a single line into an Request
func ParseRequestLine(line string) (*Request, error) {
    parts := strings.Split(line, "-")
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid format")
    }
    prefix := strings.Join(parts[:len(parts)-2], "-")
    from := parts[len(parts)-2]
    to := parts[len(parts)-1]
    return &Request{
        Prefix: prefix,
        From:   from,
        To:     to,
        Raw:    line,
    }, nil
}

// ReadRequests reads all valid requests from the given file
func ReadRequests(filename string) ([]*Request, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var requests []*Request
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip empty lines or comments
		}
		req, err := ParseRequestLine(line)
		if err != nil {
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}
		requests = append(requests, req)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return requests, nil
}
