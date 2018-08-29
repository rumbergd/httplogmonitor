package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpLogEntry_ToString(t *testing.T) {
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
	expected := "127.0.0.1 abc frank [2017-10-22:10:19:02.041890273 -0400] \"GET /section1/subsection1/1.html HTTP1.0\" 200 1234"
	actual := e.ToString()
	assert.Equal(t, actual, expected, "ToString() provided unexpected result")
}

func TestHttpLogEntry_GetWebsiteSection(t *testing.T) {
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
	expected := "section1"
	actual := e.GetWebsiteSection()
	assert.Equal(t, actual, expected, "GetWebsiteSectionring() provided unexpected result")
}

func TestHttpLogEntry_IsSuccessWith200ResponseCode(t *testing.T) {
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
	expected := true
	actual := e.IsSuccess()
	assert.Equal(t, actual, expected, "IsSuccess() provided unexpected result")
}

func TestHttpLogEntry_IsSuccessWith404ResponseCode(t *testing.T) {
	e := HttpLogEntry{
		clientIP:        "127.0.0.1",
		userIdentifier:  "abc",
		userID:          "frank",
		timestamp:       "2017-10-22:10:19:02.041890273",
		timezone:        "-0400",
		requestMethod:   "GET",
		requestResource: "/section1/subsection1/1.html",
		requestProtocol: "HTTP1.0",
		responseCode:    404,
		responseSize:    1234,
	}
	expected := false
	actual := e.IsSuccess()
	assert.Equal(t, actual, expected, "IsSuccess() provided unexpected result")
}

func TestHttpLogEntry_IsSuccessWith500ResponseCode(t *testing.T) {
	e := HttpLogEntry{
		clientIP:        "127.0.0.1",
		userIdentifier:  "abc",
		userID:          "frank",
		timestamp:       "2017-10-22:10:19:02.041890273",
		timezone:        "-0400",
		requestMethod:   "GET",
		requestResource: "/section1/subsection1/1.html",
		requestProtocol: "HTTP1.0",
		responseCode:    500,
		responseSize:    1234,
	}
	expected := false
	actual := e.IsSuccess()
	assert.Equal(t, actual, expected, "IsSuccess() provided unexpected result")
}

func TestHttpLogEntry_CreateHTTPLogEntry(t *testing.T) {
	logLine := "127.0.0.1 abc frank [2017-10-22:10:19:02.041890273 -0400] \"GET /section1/subsection1/1.html HTTP1.0\" 200 1234"
	expected := HttpLogEntry{
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

	actual := CreateHTTPLogEntry(logLine)
	assert.Equal(t, actual.clientIP, expected.clientIP, "CreateHTTPLogEntry() provided unexpected result for clientIP field")
	assert.Equal(t, actual.userIdentifier, expected.userIdentifier, "CreateHTTPLogEntry() provided unexpected result for userIdentifier field")
	assert.Equal(t, actual.userID, expected.userID, "CreateHTTPLogEntry() provided unexpected result for userID field")
	assert.Equal(t, actual.timestamp, expected.timestamp, "CreateHTTPLogEntry() provided unexpected result for timestamp field")
	assert.Equal(t, actual.timezone, expected.timezone, "CreateHTTPLogEntry() provided unexpected result for timezone field")
	assert.Equal(t, actual.requestMethod, expected.requestMethod, "CreateHTTPLogEntry() provided unexpected result for requestMethod field")
	assert.Equal(t, actual.requestResource, expected.requestResource, "CreateHTTPLogEntry() provided unexpected result for requestResource field")
	assert.Equal(t, actual.requestProtocol, expected.requestProtocol, "CreateHTTPLogEntry() provided unexpected result for requestProtocol field")
	assert.Equal(t, actual.responseCode, expected.responseCode, "CreateHTTPLogEntry() provided unexpected result for responseCode field")
	assert.Equal(t, actual.responseSize, expected.responseSize, "CreateHTTPLogEntry() provided unexpected result for responseSize field")
}
