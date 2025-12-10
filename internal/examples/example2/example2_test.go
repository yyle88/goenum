package example2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example2/internal/enums"
)

func TestEnums(t *testing.T) {
	t.Log(enums.Action)
	t.Log(enums.Action.Start())
	t.Log(enums.Action.Close())
	t.Log(enums.Action.Enums())

	_, ok := enums.Action.Enums().Lookup(enums.Action.Start())
	require.True(t, ok)

	_, ok = enums.Action.Enums().Lookup(enums.Action.Close())
	require.True(t, ok)

	_, ok = enums.Action.Enums().Lookup(enums.Action.Pause())
	require.True(t, ok)

	_, ok = enums.Action.Enums().Lookup(enums.Action.Reset())
	require.True(t, ok)

	codes := enums.Action.Enums().List()
	require.Len(t, codes, 4)
}
