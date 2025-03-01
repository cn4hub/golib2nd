// Contains a function to return a string
package echopkg

import (
	"fmt"
)

func Echo(message string) string {
	return fmt.Sprintf("Echo: ", message)
}
