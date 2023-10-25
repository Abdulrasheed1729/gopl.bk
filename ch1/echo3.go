package ch1

import "strings"

func Echo3(args []string) string {
	return strings.Join(args[1:], " ")
}
