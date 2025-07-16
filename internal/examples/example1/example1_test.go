package example1_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum"
	"github.com/yyle88/goenum/internal/examples/example1/internal/enums"
)

func TestEnums(t *testing.T) {
	t.Log(enums.Status)
	t.Log(enums.Status.OK())
	t.Log(enums.Status.WA())
	t.Log(enums.Status.Enums())

	require.True(t, slices.Contains(enums.Status.Enums(), enums.Status.OK()))
	require.True(t, slices.Contains(enums.Status.Enums(), enums.Status.WA()))

	require.True(t, goenum.Valid(enums.Status.OK()))
	require.True(t, goenum.Valid(enums.Status.WA()))
}
