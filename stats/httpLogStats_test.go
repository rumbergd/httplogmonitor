package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpLogStats_ToString(t *testing.T) {
	s := HttpLogStats{
		hits:     50,
		errors:   25,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	expected := "total hits: 50; most popular section: \"section1\" with 25 hits; average error rate: 50%"
	actual := s.ToString()
	assert.Equal(t, actual, expected, "ToString() provided unexpected result")
}

func TestHttpLogStats_GetHits(t *testing.T) {
	s := HttpLogStats{
		hits:     50,
		errors:   25,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	expected := int64(50)
	actual := s.GetHits()
	assert.Equal(t, actual, expected, "GetHits() provided unexpected result")
}

func TestHttpLogStats_GetErrorRate(t *testing.T) {
	s := HttpLogStats{
		hits:     50,
		errors:   10,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	expected := int64(20)
	actual := s.getErrorRate()
	assert.Equal(t, actual, expected, "getErrorRate() provided unexpected result")
}

func TestHttpLogStats_GetErrorRateWith0Hits(t *testing.T) {
	s := HttpLogStats{
		hits:     0,
		errors:   0,
		sections: map[string]int64{},
	}
	expected := int64(0)
	actual := s.getErrorRate()
	assert.Equal(t, actual, expected, "getErrorRate() provided unexpected result")
}

func TestHttpLogStats_GetMostPopularSection(t *testing.T) {
	s := HttpLogStats{
		hits:     50,
		errors:   10,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	expected := "section1"
	expectedNum := int64(25)
	actual, actualNum := s.getMostPopularSection()
	assert.Equal(t, actual, expected, "getMostPopularSection() provided unexpected result")
	assert.Equal(t, actualNum, expectedNum, "getMostPopularSection() provided unexpected result")
}

func TestHttpLogStats_GetMostPopularSectionWithEmptyMap(t *testing.T) {
	s := HttpLogStats{
		hits:     0,
		errors:   0,
		sections: map[string]int64{},
	}
	expected := ""
	expectedNum := int64(0)
	actual, actualNum := s.getMostPopularSection()
	assert.Equal(t, actual, expected, "getMostPopularSection() provided unexpected result")
	assert.Equal(t, actualNum, expectedNum, "getMostPopularSection() provided unexpected result")
}

func TestHttpLogStats_Reset(t *testing.T) {
	actual := HttpLogStats{
		hits:     50,
		errors:   10,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	expected := HttpLogStats{
		hits:     0,
		errors:   0,
		sections: map[string]int64{},
	}
	actual.Reset()

	assert.Equal(t, actual.hits, expected.hits, "Reset() provided unexpected result")
	assert.Equal(t, actual.errors, expected.errors, "Reset() provided unexpected result")
	assert.Equal(t, len(actual.sections), 0, "Reset() provided unexpected result")
}

func TestHttpLogStats_Add(t *testing.T) {
	actual := HttpLogStats{
		hits:     50,
		errors:   10,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	e := HttpLogEntry{
		clientIP:        "127.0.0.1",
		userIdentifier:  "abc",
		userID:          "frank",
		timestamp:       "2017-10-22:10:19:02.041890273",
		timezone:        "-0400",
		requestMethod:   "GET",
		requestResource: "/section1/subsection1/1.html",
		requestProtocol: "HTTP1.0",
		responseCode:    200,
		responseSize:    1234,
	}

	expected := HttpLogStats{
		hits:     51,
		errors:   10,
		sections: map[string]int64{"section1": 26, "section2": 15, "section3": 10},
	}
	actual.Add(e)

	assert.Equal(t, actual.hits, expected.hits, "Add() provided unexpected result")
	assert.Equal(t, actual.errors, expected.errors, "Add() provided unexpected result")
	assert.Equal(t, actual.sections, expected.sections, "Add() provided unexpected result")
}

func TestHttpLogStats_AddError(t *testing.T) {
	actual := HttpLogStats{
		hits:     50,
		errors:   10,
		sections: map[string]int64{"section1": 25, "section2": 15, "section3": 10},
	}
	e := HttpLogEntry{
		clientIP:        "127.0.0.1",
		userIdentifier:  "abc",
		userID:          "frank",
		timestamp:       "2017-10-22:10:19:02.041890273",
		timezone:        "-0400",
		requestMethod:   "GET",
		requestResource: "/section2/subsection1/1.html",
		requestProtocol: "HTTP1.0",
		responseCode:    404,
		responseSize:    1234,
	}

	expected := HttpLogStats{
		hits:     51,
		errors:   11,
		sections: map[string]int64{"section1": 25, "section2": 16, "section3": 10},
	}
	actual.Add(e)

	assert.Equal(t, actual.hits, expected.hits, "Add() provided unexpected result")
	assert.Equal(t, actual.errors, expected.errors, "Add() provided unexpected result")
	assert.Equal(t, actual.sections, expected.sections, "Add() provided unexpected result")
}
