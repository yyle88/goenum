package example5_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example5/internal/enums"
)

func TestPermissionScenario(t *testing.T) {
	require.Equal(t, "權限已開啟", processPermission("權限-開啟"))
	require.Equal(t, "權限已關閉", processPermission("權限-關閉"))
	require.Equal(t, "ERROR: INVALID PERMISSION STATE '權限-無效'", processPermission("權限-無效"))
}

// processPermission demonstrates real-world usage pattern
// processPermission 演示真實世界的使用模式
func processPermission(input string) string {
	// Step 1: Convert string to enum // 第一步：將字符串轉換為枚舉
	permissionValue := enums.PermissionEnum(input)

	// Step 2: Validate enum // 第二步：驗證枚舉合法性
	if _, ok := enums.Permission.Enums().Lookup(permissionValue); !ok {
		return fmt.Sprintf("ERROR: INVALID PERMISSION STATE '%s'", input)
	}

	// Step 3: Use in switch statement // 第三步：在 switch 語句中使用
	switch permissionValue {
	case enums.Permission.E開啟():
		return "權限已開啟"
	case enums.Permission.D關閉():
		return "權限已關閉"
	default:
		return "未知狀態"
	}
}
