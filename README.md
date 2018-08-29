# **HTTP LOG MONITOR**

## Build and Run
*This project requires Go to be installed and set up correctly. See [here](https://golang.org/doc/code.html) for more details*.

Clone this code repository into `$GOPATH/src/httplogmonitor` on your machine.
Build the project by running the following command in the project folder:

`go build` 

Adjust configuration if needed and run the executable:

`./httplogmonitor`

The program analyzes the log file continuously until you stop it with Ctrl-C.

## Configuration
HTTP Log Monitor requires proper configuration in `appconfig.yaml`.
The configuration goes as following:
- INPUTFILE (string / path to the http access log file)
- STATS (configuration section for collecting log statistics)
- - INTERVALSECONDS (integer / interval in seconds at which log statistics are gathered)
- ALERTS (configuration section for creating alerts)
- - VOLUMEALERT (configuration section for volume alerts)
- - - STATSINTERVALLOOKBACK (integer / how many log stats intervals are considered for checking traffic volume against thresholds)
- - - THRESHOLD (integer / threshold for number of hits to generate high volume alert)
- - - LOGFILE (string / file to save alerts for historical purposes)

Default configuration is:
```yaml
INPUTFILE: "http.log"
STATS:
  INTERVALSECONDS: 10
ALERTS:
  VOLUMEALERT: 
    STATSINTERVALLOOKBACK: 12 # to evaluate traffic for the last 2 min (12 * 10 sec)
    THRESHOLD: 100
    LOGFILE: "volume_alerts.log"
```

## Unit tests

Run unit tests from the project folder using the command:

`go test -v ./...`


## Dependency management
External dependencies/libraries are managed by *[govendor](https://github.com/kardianos/govendor)*.


All external dependencies are specified in this file in the project folder:
`vendor/vendor.json`

## Design

### Workflow

HTTP Log Monitor loads the configuration from `appconfig.yaml` during startup. 

Then it starts a timer for collecting the statistics about the http access log. 

Then it starts tailing the http access log file. 

### Log format

Lines in the http access log must be in Common Log Format - https://en.wikipedia.org/wiki/Common_Log_Format

Some examples:
```
64.242.88.16 - - [25/Oct/2017:10:02:42 -0400] "GET /pictures/5.gif HTTP/1.0" 200 7645
64.242.88.19 - - [25/Oct/2017:10:02:44 -0400] "GET /slides/1.html HTTP/1.0" 200 9845
64.242.88.10 - - [25/Oct/2017:10:02:46 -0400] "GET /pages/1.html HTTP/1.0" 200 2326
```


### Statistics Collection

Log statistics are updated for every new line appended to the http access log. Log stats collection is cumulative.

At every timer interval (10 sec by default), the HTTP Log Moitor prints statitics collected for that period. Log stats are only printed to console.

Currently the log statistics print the following:
- total number of hits
- most popular website section with its number of hits
- average error rate

### Alerts

Statistics are evaluated for generating an alert at every timer interval (10 sec by default). 

However, the data collected for the alert is more than just the last timer interval. Number of timer intervals to look back is determined by `STATSINTERVALLOOKBACK` configuration setting. By default, it's 12 *(12 * 10 sec = 2 min)*.

Currently, the only alert type supported is a high http traffic volume alert. Default threshold for this alert is 100 hits per 2 minute. 

Alerts are printed to console and saved into a file for historical reference.

Tail the alerts log file using `tail -f <alerts_log_file>` command to monitor the alerts as the monitor runs.

## Possible Improvements
- Collect more useful statistics from http access log. For example,
- - Show resources with 404 errors and clients requesting them
- - Show not just average error rate, but a rate of 404s, or rate of 5XX errors to understand what might be wrong with the website
- Add more alert types. For example,
- - Error rate over a certain threshold alert 
