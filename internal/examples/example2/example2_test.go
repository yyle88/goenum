package example2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum"
	"github.com/yyle88/goenum/internal/examples/example2/internal/enums"
)

func TestEnums(t *testing.T) {
	t.Log(enums.Action)
	t.Log(enums.Action.Start())
	t.Log(enums.Action.Close())
	t.Log(enums.Action.Enums())

	require.True(t, slices.Contains(enums.Action.Enums(), enums.Action.Start()))
	require.True(t, slices.Contains(enums.Action.Enums(), enums.Action.Close()))

	require.True(t, goenum.Valid(enums.Action.Start()))
	require.True(t, goenum.Valid(enums.Action.Close()))
	require.True(t, goenum.Check(enums.Action))
	require.True(t, goenum.Check(enums.Action.Start()))
	require.True(t, goenum.Check(enums.Action.Close()))

	require.True(t, enums.Action.Check())
	require.True(t, enums.Action.Start().Valid())
	require.True(t, enums.Action.Close().Valid())
	require.True(t, enums.Action.Pause().Check())
	require.True(t, enums.Action.Reset().Check())
}
