package gitwatch

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xaque208/znet/pkg/events"
)

func TestProducer_interface(t *testing.T) {
	var a events.Producer = &EventProducer{}
	require.NotNil(t, a)
}
