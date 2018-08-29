package stats

import (
	"fmt"
	"strings"
)

// HttpLogEntry is a type representing line of the http log format according to Common Log Format - https://en.wikipedia.org/wiki/Common_Log_Format
type HttpLogEntry struct {
	clientIP        string
	userIdentifier  string
	userID          string
	timestamp       string
	timezone        string
	requestMethod   string
	requestResource string
	requestProtocol string
	responseCode    int
	responseSize    int
}

// ToString returns a string representation of HttpLogEntry
func (e *HttpLogEntry) ToString() string {
	return fmt.Sprintf("%s %s %s [%s %s] \"%s %s %s\" %d %d",
		e.clientIP, e.userIdentifier, e.userID,
		e.timestamp, e.timezone,
		e.requestMethod, e.requestResource, e.requestProtocol,
		e.responseCode, e.responseSize,
	)
}

// GetWebsiteSection returns substring of requestResource field between 1st and 2nd slashes
func (e *HttpLogEntry) GetWebsiteSection() string {
	s := strings.Split(e.requestResource, "/")
	if len(s) > 1 {
		return s[1]
	}
	return ""
}

// IsSuccess returns true if responseCode field is 2XX
func (e *HttpLogEntry) IsSuccess() bool {
	firstDigit := e.responseCode / 100
	if firstDigit == 2 {
		return true
	}
	return false
}

// CreateHTTPLogEntry creates HttpLogEntry from a log line
func CreateHTTPLogEntry(logLine string) HttpLogEntry {
	var e HttpLogEntry
	logLine = strings.Replace(logLine, "[", "", 1)
	logLine = strings.Replace(logLine, "]", "", 1)
	logLine = strings.Replace(logLine, "\"", "", 2)

	fmt.Sscanf(logLine, "%s %s %s %s %s %s %s %s %d %d",
		&e.clientIP, &e.userIdentifier, &e.userID,
		&e.timestamp, &e.timezone,
		&e.requestMethod, &e.requestResource, &e.requestProtocol,
		&e.responseCode, &e.responseSize,
	)
	return e
}
