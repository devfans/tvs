package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Channel represents a single channel in the M3U file.
type Channel struct {
	Attributes map[string]string
	Title      string
	Duration   int
	Url        string
}

// Group returns the group-title attribute of the channel.
func (c *Channel) Group() string {
	return c.Attributes["group-title"]
}

// Source represents the entire M3U file.
type Source struct {
	Url      string
	Channels map[string]*Channel
}

// parseM3UFile reads an M3U file and parses it into a Source struct.
func parseM3UFile(filePath string) (*Source, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	source := &Source{
		Channels: make(map[string]*Channel),
	}

	var currentChannel *Channel

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#EXTINF:") {
			// Parse the EXTINF line to get duration, title, and attributes
			parts := strings.SplitN(line[len("#EXTINF:"):], ",", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid EXTINF line: %s", line)
			}

			duration, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid duration in EXTINF line: %s", line)
			}

			title := parts[1]
			attributes := parseAttributes(line)

			currentChannel = &Channel{
				Attributes: attributes,
				Title:      title,
				Duration:   duration,
			}
		} else if !strings.HasPrefix(line, "#") && len(line) > 0 {
			// This line is the URL for the current channel
			if currentChannel == nil {
				return nil, fmt.Errorf("URL found without preceding EXTINF line: %s", line)
			}

			currentChannel.Url = line
			source.Channels[currentChannel.Title] = currentChannel
			currentChannel = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return source, nil
}

// parseAttributes extracts key-value pairs from an EXTINF line.
func parseAttributes(line string) map[string]string {
	attributes := make(map[string]string)
	start := strings.Index(line, " ")
	if start == -1 {
		return attributes
	}

	parts := strings.Split(line[start+1:], " ")
	for _, part := range parts {
		if strings.Contains(part, "=") {
			kv := strings.SplitN(part, "=", 2)
			if len(kv) == 2 {
				attributes[strings.Trim(kv[0], `"`)] = strings.Trim(kv[1], `"`)
			}
		}
	}

	return attributes
}

func main() {
	// Example usage
	source, err := parseM3UFile("example.m3u")
	if err != nil {
		fmt.Printf("Error parsing M3U file: %v\n", err)
		return
	}

	for title, channel := range source.Channels {
		fmt.Printf("Channel: %s\n", title)
		fmt.Printf("  Group: %s\n", channel.Group())
		fmt.Printf("  Duration: %d\n", channel.Duration)
		fmt.Printf("  URL: %s\n", channel.Url)
		fmt.Println()
	}
}
