package logmessage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestBuildMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "C", "D")
	fmt.Println(actual)
	assert.NotEmpty(test, actual)
}

func TestBuildMessageNoDetails(test *testing.T) {
	actual := BuildMessage("A", "B", "C")
	fmt.Println(actual)
}

func TestBuildMessageNoMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "")
	fmt.Println(actual)
}

func TestBuildMessageNoLevel(test *testing.T) {
	actual := BuildMessage("A", "", "")
	fmt.Println(actual)
}

func TestBuildMessageUsingMap(test *testing.T) {
	detailsMap := map[string]string{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
	}
	actual := BuildMessageUsingMap("A", "B", "C", detailsMap)
	fmt.Println(actual)
}

func TestParseMessage(test *testing.T) {
	message := BuildMessage("A", "B", "C", "D")
	parsedMessage := ParseMessage(message)

	fmt.Println(parsedMessage.Level)
}

func TestParseMessageUsingMap(test *testing.T) {
	detailsMap := map[string]string{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
	}
	message := BuildMessageUsingMap("A", "B", "C", detailsMap)
	parsedMessage := ParseMessage(message)
	details, ok := parsedMessage.Details.(map[string]interface{})
	if !ok {
		fmt.Printf("Error: %t", ok)
	}

	fmt.Println(details["FirstVariable"])
}
