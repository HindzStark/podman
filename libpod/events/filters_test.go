//go:build linux || freebsd

package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateEventFilterNetwork(t *testing.T) {
	filterFunc, err := generateEventFilter("network", "mynet")
	require.NoError(t, err)

	netEventMatch := &Event{
		Type: Network,
		Name: "mynet",
		ID:   "1234567890abcdef",
	}
	netEventMismatch := &Event{
		Type: Network,
		Name: "othernet",
		ID:   "9876543210fedcba",
	}
	nonNetEvent := &Event{
		Type: Container,
		Name: "mynet",
		ID:   "1234567890abcdef",
	}

	assert.True(t, filterFunc(netEventMatch))
	assert.False(t, filterFunc(netEventMismatch))
	assert.False(t, filterFunc(nonNetEvent))

	// Test prefix ID match for network filter
	filterFuncID, err := generateEventFilter("network", "12345")
	require.NoError(t, err)
	assert.True(t, filterFuncID(netEventMatch))
}
