package main

import (
	"strconv"
	"strings"
)

func decodeMessage(msg string, l int) []string {
	if msg[0] == '+' || msg[0] == '-' || msg[0] == ':' {
		return []string{msg[1 : l-2]}
	} else if msg[0] == '$' {
		msgs := strings.Split(msg, "\r\n")
		if len(msgs) < 2 {
			return []string{"response error: $ type size error"}
		}

		size, err := strconv.Atoi(msgs[0][1:])
		if err != nil {
			return []string{"response error: $ type size invalid int"}
		}

		return []string{msgs[1][:size]}
	} else {
		return []string{}
	}
}
