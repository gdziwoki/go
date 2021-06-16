package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackage(t *testing.T) {
	modulePackage, err := Package("../..")
	require.NoError(t, err)
	assert.Equal(t, "github.com/gdziwoki/go/gopages", modulePackage)
}
