package client

import (
	"strings"
)

func sanitizeIdentifier(identifier string) string {
	identifier = strings.ReplaceAll(identifier, string([]byte{0}), "")

	if strings.HasPrefix(identifier, `[`) && strings.HasSuffix(identifier, `]`) {
		return identifier
	}
	return `[` + identifier + `]`
}
