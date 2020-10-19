// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Copied from sirupsen/logrus, so that LICENSE applies to this file.

package log

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	// DefaultTimestampFormat is the time format with milliseconds.
	DefaultTimestampFormat string = "2006-01-02T15:04:05.999Z07:00"

	// MaxMessageBodySize is the max length of the log message before it is trimmed.
	MaxMessageBodySize int = 1 << 14
)

var TimestampFormat = DefaultTimestampFormat

// LogstashFormatter generates json in logstash format.
// See the Logstash website http://logstash.net/
type LogstashFormatter struct {
	Type     string // if not empty use for logstash type field.
	Hostname string
}

func trimMessage(message string) string {
	if len(message) > MaxMessageBodySize {
		message = message[:MaxMessageBodySize] + " ... (message has been trimmed due to extreme length)"
	}
	return message
}

// We can not modify entry.Data as it is shared between multiple threads
func (f *LogstashFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	dataCopy := map[string]interface{}{}
	for key, val := range entry.Data {
		dataCopy[key] = val
	}
	dataCopy["@version"] = 1
	dataCopy["HOSTNAME"] = f.Hostname
	dataCopy["@timestamp"] = entry.Time.UTC().Format(TimestampFormat)

	// set message field
	message := entry.Message
	if dataCopy[ErrorKey] != nil {
		// When searching logs for a message, users commonly also intend to search the error
		// text. As it is not clear how to make this behaviour the default in Kibana, we
		// instead ensure all remote logs have the error description in the message.
		// TODO(IC-355): Fix Kibana so it searches all relevant fields by default.
		message = fmt.Sprintf("%v: %v", message, dataCopy["error"])
	}
	dataCopy["message"] = trimMessage(message)

	// Error is also a candidate for being large
	if dataCopy[ErrorKey] != nil {
		dataCopy[ErrorKey] = trimMessage(fmt.Sprintf("%v", dataCopy[ErrorKey]))
	}
	if dataCopy[ErrorStackKey] != nil {
		dataCopy[ErrorStackKey] = trimMessage(fmt.Sprintf("%v", dataCopy[ErrorStackKey]))
	}

	level := strings.ToUpper(entry.Level.String())
	// we assume that warning has level WARN inside kibana
	if level == "WARNING" {
		level = "WARN"
	}
	// set level field
	dataCopy["level"] = level

	// set type field
	if f.Type != "" {
		dataCopy["type"] = f.Type
	}

	serialized, err := json.Marshal(dataCopy)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
