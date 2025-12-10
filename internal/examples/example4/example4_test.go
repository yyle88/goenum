package example4_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example4/internal/enums"
)

func TestStatusScenario(t *testing.T) {
	require.Equal(t, "任务成功", processStatus("状态-成功"))
	require.Equal(t, "任务失败", processStatus("状态-失败"))
	require.Equal(t, "ERROR: INVALID STATUS '状态-取消'", processStatus("状态-取消"))
}

// processStatus demonstrates Suffix naming mode usage
// processStatus 演示 Suffix 命名模式的使用
func processStatus(input string) string {
	// Step 1: Convert string to enum // 第一步：将字符串转换为枚举
	statusValue := enums.StatusEnum(input)

	// Step 2: Validate enum // 第二步：验证枚举合法性
	if _, ok := enums.Status.Enums().Lookup(statusValue); !ok {
		return fmt.Sprintf("ERROR: INVALID STATUS '%s'", input)
	}

	// Step 3: Use in switch statement // 第三步：在 switch 语句中使用
	switch statusValue {
	case enums.Status.S成功():
		return "任务成功"
	case enums.Status.F失败():
		return "任务失败"
	case enums.Status.P等待():
		return "任务等待"
	default:
		return "未知状态"
	}
}
