// Contains a function to return a string
package echo_test

import (
	"fmt"
)

func Echo(message string) string {
	return fmt.Sprintf("Echo: ", message)
}
