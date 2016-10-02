package gokk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRFC3339Regexp(t *testing.T) {
	const prefix = "zfs-auto-snap"

	for _, tt := range []struct {
		matchExpected bool
		s             string
	}{
		// From the examples section of the RFC itself.
		{true, "1985-04-12T23:20:50.52Z"},
		{true, "1996-12-19T16:39:57-08:00"},
		{true, "1990-12-31T23:59:60Z"},
		{true, "1990-12-31T15:59:60-08:00"},
		{true, "1937-01-01T12:00:27.87+00:20"},
	} {
		m := RFC3339Regexp.FindStringSubmatch(tt.s)
		matched := len(m) > 0

		if tt.matchExpected {
			assert.True(t, matched, fmt.Sprintf("expected input to match regexp: %q", tt.s))
		} else {
			assert.False(t, matched, fmt.Sprintf("did not expect input to match regexp: %q", tt.s))
		}
	}
}
