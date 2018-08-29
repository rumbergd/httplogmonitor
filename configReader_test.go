package main

import (
	. "gopkg.in/check.v1"
	"httplogmonitor/config"
	"os"
	"strconv"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestPackageMain(t *testing.T) { TestingT(t) }

// Test suite structure for individual config tests
type configTestSuite struct {
	testNum   int
	dirConfig string
}

var _ = Suite(&configTestSuite{})

//Setup Test Suite
func (s *configTestSuite) SetUpSuite(c *C) {
	s.dirConfig = c.MkDir()
}

//Setup run before each individual test
func (s *configTestSuite) SetUpTest(c *C) {
	s.testNum++
}

// Test loading of the INPUTFILE
func (s *configTestSuite) TestConfigInputFile(c *C) {
	configString := string("INPUTFILE: \"test.log\"")
	testFileName := createConfigTestFile(s, configString)
	loadConfigurations(testFileName)

	c.Check(config.InputFile, Equals, "test.log")
}

// Test loading of the STATS section
func (s *configTestSuite) TestConfigStatsSection(c *C) {
	configString := string("STATS:\r\n INTERVALSECONDS: 10")
	testFileName := createConfigTestFile(s, configString)
	loadConfigurations(testFileName)

	c.Check(config.StatsConfig.IntervalSeconds, Equals, 10)
}

// Test loading of the ALERTS section
func (s *configTestSuite) TestConfigAlertsSection(c *C) {
	configString := string("ALERTS:\r\n VOLUMEALERT:\r\n  STATSINTERVALLOOKBACK: 12\r\n  THRESHOLD: 100\r\n  LOGFILE: \"alerts.log\"")
	testFileName := createConfigTestFile(s, configString)
	loadConfigurations(testFileName)

	c.Check(config.AlertsConfig.VolumeAlert.StatsIntervalLookback, Equals, 12)
	c.Check(config.AlertsConfig.VolumeAlert.Threshold, Equals, int64(100))
	c.Check(config.AlertsConfig.VolumeAlert.LogFile, Equals, "alerts.log")
}

// Creates a temp config file in a temp directory for one of the unit tests
func createConfigTestFile(s *configTestSuite, buffer string) string {

	filename := s.dirConfig + "/test" + strconv.Itoa(s.testNum) + ".yaml"
	f, e := os.Create(filename)
	check(e)
	defer f.Close()

	_, e = f.WriteString(buffer)
	check(e)
	return filename
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
