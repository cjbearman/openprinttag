package vtag

import (
	"fmt"
	"os"
)

// DebugMode can be set to true for debug logging, which is really useful for development
var DebugMode = false

// debug is called by the code to write debug messages when DebugMode is enabled
func debug(format string, v ...any) {
	if DebugMode {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("debug: %s\n", format), v...)
	}
}
