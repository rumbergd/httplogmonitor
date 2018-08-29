package stats

import (
	"fmt"
	"sync"
	"time"
)

// LogStats represents statistics about http traffic collected at timer intervals
var LogStats HttpLogStats

// HttpLogStats is a type representing stats collected about the http traffic on certain intervals
type HttpLogStats struct {
	sync.RWMutex
	hits     int64
	errors   int64
	sections map[string]int64
}

// Add updates HttpLogStats with a new http log entry
func (s *HttpLogStats) Add(e HttpLogEntry) {
	// Ensure the access is thread safe
	s.Lock()
	defer s.Unlock()

	// Add hits
	s.hits++

	// Add errors
	if !e.IsSuccess() {
		s.errors++
	}

	// Add sections
	if s.sections == nil {
		s.sections = make(map[string]int64)
	}
	s.sections[e.GetWebsiteSection()]++
}

// Reset zeroes all the collected statistics up to this moment
func (s *HttpLogStats) Reset() {
	// Ensure the access is thread safe
	s.Lock()
	defer s.Unlock()

	s.hits = 0
	s.errors = 0
	s.sections = make(map[string]int64)
}

// ToString returns a string representation of HttpLogStats
func (s *HttpLogStats) ToString() string {
	// Ensure the access is thread safe
	s.Lock()
	defer s.Unlock()

	mostPopularSection, mostPopularSectionHits := s.getMostPopularSection()
	return fmt.Sprintf("total hits: %v; most popular section: \"%v\" with %v hits; average error rate: %v%%",
		s.hits, mostPopularSection, mostPopularSectionHits, s.getErrorRate(),
	)
}

// Log uses logger to record stats
func (s *HttpLogStats) Log() {
	statsLogger.Printf("[%v] HTTP log statistics for the last 10 sec / %v\n", time.Now(), s.ToString())
}

// GetHits returns hits field from the HttpLogStats
func (s *HttpLogStats) GetHits() int64 {
	// Ensure the access is thread safe
	s.Lock()
	defer s.Unlock()

	return s.hits
}

// getMostPopularSection traverses map of sections and returns the one with most hits
func (s *HttpLogStats) getMostPopularSection() (string, int64) {

	mostPopularSection := ""
	mostPopularSectionHits := int64(0)

	if s.sections != nil {
		for section, hits := range s.sections {
			if hits > mostPopularSectionHits {
				mostPopularSection = section
				mostPopularSectionHits = hits
			}
		}
	}

	return mostPopularSection, mostPopularSectionHits
}

// getErrorRate computes error rate
func (s *HttpLogStats) getErrorRate() int64 {
	if s.hits > 0 {
		return s.errors * 100 / s.hits
	}
	return 0
}
