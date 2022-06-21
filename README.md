# go-json-log-message

## Synopsis

A simple Go package that formats (error) messages as JSON.
When pretty printed, the messages have the following format:

```json
{
    "id": "a-unique-message-id",
    "text": "Explanation of the error",
    "details": {
        "firstVariable": "First value",
        "secondVariable": "Second value"
    }
}
```

## Usage

To import the go package:

```go
import (
    :
    "github.com/docktermj/go-json-log-message/message"
)
```

For examples of use, see:

1. [main.go](main.go)
1. [message_test.go](message/message_test.go)


## Development

1. To run unit test cases.
   Example:

    ```console
    make test
    ```

1. To build `main.go`.
   Example:

    ```console
    make build
    ```

   The output file will be at `target/linux/go-json-log-message`.


1. To clean up.
   Example:

    ```console
    make clean
    ```

