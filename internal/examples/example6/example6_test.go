package example6_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example6/internal/enums"
)

func TestConnectionScenario(t *testing.T) {
	require.Equal(t, "接続が確立されました", processConnection("接続接続"))
	require.Equal(t, "接続が切断されました", processConnection("接続切断"))
	require.Equal(t, "接続待機中です", processConnection("接続待機"))
	require.Equal(t, "ERROR: INVALID CONNECTION STATE '接続無効'", processConnection("接続無効"))
}

// processConnection demonstrates real-world usage pattern
// processConnection 現実世界の使用パターンを実演
func processConnection(input string) string {
	// Step 1: Convert string to enum // ステップ1：文字列を列挙型に変換
	connectionValue := enums.ConnectionEnum(input)

	// Step 2: Validate enum // ステップ2：列挙型の検証
	if _, ok := enums.Connection.Enums().Lookup(connectionValue); !ok {
		return fmt.Sprintf("ERROR: INVALID CONNECTION STATE '%s'", input)
	}

	// Step 3: Use in switch statement // ステップ3：switch文で使用
	switch connectionValue {
	case enums.Connection.C接続():
		return "接続が確立されました"
	case enums.Connection.D切断():
		return "接続が切断されました"
	case enums.Connection.W待機():
		return "接続待機中です"
	default:
		return "不明な状態"
	}
}
