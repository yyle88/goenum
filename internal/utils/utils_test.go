package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/utils"
)

func TestCodeString(t *testing.T) {
	require.Equal(t, "0", utils.CodeString(0))
	require.Equal(t, "true", utils.CodeString(true))
	require.Equal(t, `"abc"`, utils.CodeString("abc"))
}

func TestWithQuotes(t *testing.T) {
	a := `"` + "abc" + `"`
	b := utils.TrimQuotes(a)
	require.Equal(t, "abc", b)
	c := utils.WithQuotes(b)
	require.Equal(t, a, c)
}
