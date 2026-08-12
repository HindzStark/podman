//go:build linux || freebsd

package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateEventFilterMachine(t *testing.T) {
	t.Parallel()

	filterFunc, err := generateEventFilter("machine", "podman-machine-default")
	require.NoError(t, err)
	require.NotNil(t, filterFunc)

	// Matching machine event by name
	machineEvent := &Event{
		Type: Machine,
		Name: "podman-machine-default",
		ID:   "1234567890abcdef",
	}
	assert.True(t, filterFunc(machineEvent))

	// Matching machine event by ID prefix
	filterFuncByID, err := generateEventFilter("machine", "12345")
	require.NoError(t, err)
	assert.True(t, filterFuncByID(machineEvent))

	// Non-matching machine event
	otherMachineEvent := &Event{
		Type: Machine,
		Name: "other-machine",
		ID:   "9876543210fedcba",
	}
	assert.False(t, filterFunc(otherMachineEvent))

	// Non-machine event type with same name
	containerEvent := &Event{
		Type: Container,
		Name: "podman-machine-default",
		ID:   "1234567890abcdef",
	}
	assert.False(t, filterFunc(containerEvent))
}
