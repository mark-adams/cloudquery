package client

import (
	"strings"
)

func sanitizeIdentifier(parts ...string) string {
	res := make([]string, 0, len(parts))

	for _, part := range parts {
		res = append(res, `[`+strings.ReplaceAll(part, string([]byte{0}), "")+`]`)
	}

	return strings.Join(parts, ".")
}
