package goenum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/goenumgen"
)

func TestEnums(t *testing.T) {
	t.Log(goenumgen.NamingMode)
	t.Log(goenumgen.NamingMode.Prefix())
	t.Log(goenumgen.NamingMode.Suffix())
	t.Log(goenumgen.NamingMode.Enums())

	_, ok := goenumgen.NamingMode.Enums().Lookup(goenumgen.NamingMode.Middle())
	require.True(t, ok)

	_, ok = goenumgen.NamingMode.Enums().Lookup(goenumgen.NamingMode.Single())
	require.True(t, ok)

	_, ok = goenumgen.NamingMode.Enums().Lookup(goenumgen.NamingMode.Suffix())
	require.True(t, ok)

	codes := goenumgen.NamingMode.Enums().List()
	require.Len(t, codes, 4)
}
