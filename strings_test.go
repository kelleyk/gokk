package gokk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	for _, tt := range []struct {
		s, sep   string
		expected []string
	}{
		{"foo bar baz", " ", []string{"foo", " ", "bar baz"}},
		{"foo  bar baz", " ", []string{"foo", " ", " bar baz"}},
		{" foo bar baz", " ", []string{"", " ", "foo bar baz"}},
		{"foobarbaz", " ", []string{"", "", "foobarbaz"}},
	} {
		r0, r1, r2 := Partition(tt.s, tt.sep)
		assert.Equal(t, tt.expected[0], r0)
		assert.Equal(t, tt.expected[1], r1)
		assert.Equal(t, tt.expected[2], r2)
	}
}

func TestPartitionLast(t *testing.T) {
	for _, tt := range []struct {
		s, sep   string
		expected []string
	}{
		{"foo bar baz", " ", []string{"foo bar", " ", "baz"}},
		{"foo bar  baz", " ", []string{"foo bar ", " ", "baz"}},
		{"foo bar baz ", " ", []string{"foo bar baz", " ", ""}},
		{"foobarbaz", " ", []string{"foobarbaz", "", ""}},
	} {
		r0, r1, r2 := PartitionLast(tt.s, tt.sep)
		assert.Equal(t, tt.expected[0], r0)
		assert.Equal(t, tt.expected[1], r1)
		assert.Equal(t, tt.expected[2], r2)
	}
}
