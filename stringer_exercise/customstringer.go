package customstringer

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
	parts := make([]string, len(p))
	for i, v := range p {
		parts[i] = fmt.Sprint(v)
	}
	return strings.Join(parts, ".")
}
