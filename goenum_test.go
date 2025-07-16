package goenum_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum"
	"github.com/yyle88/goenum/goenumgen"
)

func TestEnums(t *testing.T) {
	t.Log(goenumgen.NamingMode)
	t.Log(goenumgen.NamingMode.Prefix())
	t.Log(goenumgen.NamingMode.Suffix())
	t.Log(goenumgen.NamingMode.Enums())

	require.True(t, slices.Contains(goenumgen.NamingMode.Enums(), goenumgen.NamingMode.Middle()))
	require.True(t, slices.Contains(goenumgen.NamingMode.Enums(), goenumgen.NamingMode.Single()))

	require.True(t, goenum.Valid(goenumgen.NamingMode.Suffix()))
	require.True(t, goenum.Valid(goenumgen.NamingMode.Single()))
}
