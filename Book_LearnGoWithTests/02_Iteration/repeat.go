// Package iteration shows basic usage of iteration in golang
package iteration

import "strings"

func Repeat(str string, repeat int) string {
	var repeated strings.Builder
	for range repeat {
		repeated.WriteString(str)
	}
	return repeated.String()
}
