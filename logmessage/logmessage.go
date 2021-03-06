/*
Package message ...
*/
package logmessage

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func isJson(unknownString string) bool {
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownString), &jsonString) == nil
}

func jsonAsInterface(unknownString string) interface{} {
	var jsonString json.RawMessage
	json.Unmarshal([]byte(unknownString), &jsonString)
	return jsonString
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build a message given details as strings.
func BuildMessage(id string, level string, text string, details ...string) string {

	resultStruct := Message{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
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

// Build a message from an error
func BuildMessageFromError(id string, level string, text string, err error, details ...string) string {

	resultStruct := Message{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
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

	// Nest prior Error message.

	if err != nil {
		errorMessage := err.Error()

		var priorError interface{}
		if isJson(errorMessage) {
			priorError = jsonAsInterface(errorMessage)
		} else {
			priorError = Message{
				Text: errorMessage,
			}
		}
		resultStruct.Error = priorError
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageFromErrorUsingMap(id string, level string, text string, err error, details map[string]string) string {

	resultStruct := Message{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
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

	// Nest prior Error message.

	if err != nil {
		errorMessage := err.Error()
		var priorError interface{}
		if isJson(errorMessage) {
			priorError = jsonAsInterface(errorMessage)
		} else {
			priorError = Message{
				Text: errorMessage,
			}
		}
		resultStruct.Error = priorError
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageUsingMap(id string, level string, text string, details map[string]string) string {

	resultStruct := Message{}

	// Fill optional fields.

	if len(id) > 0 {
		resultStruct.Id = id
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

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Parse JSON message.
func ParseMessage(jsonString string) Message {
	var message Message
	json.Unmarshal([]byte(jsonString), &message)
	return message
}
