package example1_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example1/internal/enums"
)

func TestEnums(t *testing.T) {
	t.Log(enums.Status)
	t.Log(enums.Status.OK())
	t.Log(enums.Status.WA())
	t.Log(enums.Status.Enums())

	// Use enum.Enums methods for lookup and validation
	// 使用 enum.Enums 方法进行查找和验证
	_, ok := enums.Status.Enums().Lookup(enums.Status.OK())
	require.True(t, ok)

	_, ok = enums.Status.Enums().Lookup(enums.Status.WA())
	require.True(t, ok)

	// List all enum values
	// 列出所有枚举值
	codes := enums.Status.Enums().List()
	require.Len(t, codes, 2)
	require.Contains(t, codes, enums.Status.OK())
	require.Contains(t, codes, enums.Status.WA())
}
