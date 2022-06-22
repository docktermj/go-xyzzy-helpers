/*
Package message ...
*/
package logmessage

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build a message given details as strings.
func BuildMessage(id string, level string, text string, details ...string) string {

	resultStruct := Message{
		Id: id,
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		resultStruct.Text = text
	}

	// Construct details.

	detailMap := make(map[string]interface{})
	for index, value := range details {
		detailMap[strconv.Itoa(index+1)] = value
	}

	if len(details) > 0 {
		resultStruct.Details = detailMap
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageUsingMap(id string, level string, text string, details map[string]string) string {

	resultStruct := Message{
		Id: id,
	}

	if len(level) > 0 {
		resultStruct.Level = level
	}

	if len(text) > 0 {
		resultStruct.Text = text
	}

	if len(details) > 0 {
		resultStruct.Details = details
	}

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Parse JSON message.
func ParseMessage(jsonString string) Message {
	var message Message
	json.Unmarshal([]byte(jsonString), &message)
	return message
}
