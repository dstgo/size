package size

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPattern(t *testing.T) {
	assert.True(t, pattern.MatchString("1.2"))
	assert.True(t, pattern.MatchString("1.2KB"))
	assert.True(t, pattern.MatchString("1.2B"))
	assert.True(t, pattern.MatchString("1024.0091GB"))
	assert.True(t, pattern.MatchString("1024.0091g"))

	assert.False(t, pattern.MatchString("-1.2B"))
	assert.False(t, pattern.MatchString("-1.2B1123"))
	assert.False(t, pattern.MatchString(""))
}

func TestLookup(t *testing.T) {
	samples := []struct {
		data string
		size any
		ok   bool
	}{
		{data: "1.2MB", size: New(1.2, MB), ok: true},
		{data: "1B", size: New(1, B), ok: true},
		{data: "2.5GB", size: New(2.5, GB), ok: true},
		{data: "1024.111GB", size: New(1024.111, GB), ok: true},
		{data: "65535KB", size: New(65535, KB), ok: true},
		{data: "65535kb", size: New(65535, KB), ok: true},
		{data: "65535Kb", size: New(65535, KB), ok: true},
		{data: "65535kB", size: New(65535, KB), ok: true},

		{data: "-1", size: nil, ok: false},
		{data: "-&*6", size: nil, ok: false},
		{data: "-1KB", size: nil, ok: false},
		{data: "0B", size: New(0, B), ok: true},
		{data: "00000B", size: New(0, B), ok: true},
		{data: "00101B", size: New(101, B), ok: true},
	}

	for _, sample := range samples {
		size, ok := Lookup(sample.data)
		assert.EqualValues(t, sample.ok, ok)
		if ok {
			assert.EqualValues(t, sample.size, size)
		}
	}
}

func TestLookupTo(t *testing.T) {
	samples := []struct {
		data string
		to   Unit
		size any
		ok   bool
	}{
		{data: "1.2MB", to: KB, size: New(1.2*float64(MB/KB), KB), ok: true},
		{data: "1.2GB", to: KB, size: New(1.2*float64(GB/KB), KB), ok: true},
		{data: "1.2MB", to: GB, size: New(1.2*float64(MB)/float64(GB), GB), ok: true},
		{data: "-1.2GB", to: KB, size: nil, ok: false},
	}

	for _, sample := range samples {
		size, ok := LookupTo(sample.data, sample.to)
		assert.EqualValues(t, sample.ok, ok)
		if ok {
			assert.EqualValues(t, sample.size, size)
		}
	}
}

func TestSize_String(t *testing.T) {
	samples := []struct {
		data     Size
		expected string
	}{
		{data: New(1.2, KB), expected: "1.2KB"},
		{data: New(1.20000, KB), expected: "1.2KB"},
		{data: New(1.2967, KB), expected: "1.297KB"},
	}

	for _, sample := range samples {
		assert.EqualValues(t, sample.expected, sample.data.String())
	}
}
