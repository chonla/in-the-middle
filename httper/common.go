package httper

import (
    "fmt"
)

type HeaderKeyPair struct {
	Key   string
	Value string
}

func (h *HeaderKeyPair) ToString() string {
    return fmt.Sprintf("\"%s: %s\"", h.Key, h.Value)
}
