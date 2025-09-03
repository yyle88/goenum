package example3_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goenum/internal/examples/example3/internal/enums"
)

func TestSwitchScenario(t *testing.T) {
	require.Equal(t, "设备已开启", processSwitch("开启"))
	require.Equal(t, "设备已关闭", processSwitch("关闭"))
	require.Equal(t, "ERROR: INVALID SWITCH STATE '无效状态'", processSwitch("无效状态"))
}

// processSwitch demonstrates real-world usage pattern
// processSwitch 演示真实世界的使用模式
func processSwitch(input string) string {
	// Step 1: Convert string to enum // 第一步：将字符串转换为枚举
	switchValue := enums.SwitchEnum(input)

	// Step 2: Validate enum // 第二步：验证枚举合法性
	if !switchValue.Valid() {
		return fmt.Sprintf("ERROR: INVALID SWITCH STATE '%s'", input)
	}

	// Step 3: Use in switch statement // 第三步：在 switch 语句中使用
	switch switchValue {
	case enums.Switch.On():
		return "设备已开启"
	case enums.Switch.Off():
		return "设备已关闭"
	default:
		return "未知状态"
	}
}
