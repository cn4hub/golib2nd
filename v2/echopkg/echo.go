// Contains a function to return a string
package echopkg

import (
	"fmt"
)

func Echo(message string, botname string) string {
	return fmt.Sprintf("Echo (%v): %v", botname, message)
}
