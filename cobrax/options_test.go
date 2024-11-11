package cobrax

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithVersion(t *testing.T) {
	cmd := &cobra.Command{}
	WithVersion("1.0.0")(cmd)
	assert.Equal(t, "1.0.0", cmd.Version)
	require.NotNil(t, cmd.Annotations)
	assert.Equal(t, "1.0.0", cmd.Annotations[VersionKey])
}
