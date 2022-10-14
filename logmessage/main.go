// The message package formats messages into a JSON string.
package logmessage

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Message struct {
	Id       string      `json:"id,omitempty"`
	Level    string      `json:"level,omitempty"`
	Text     string      `json:"text,omitempty"`
	TextJson interface{} `json:"textJson,omitempty"`
	Details  interface{} `json:"details,omitempty"`
	Error    interface{} `json:"error,omitempty"`
}
