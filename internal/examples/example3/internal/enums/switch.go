// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: switch_test.go:11 -> enums_test.TestGenerateSwitch
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"github.com/yyle88/goenum"
)

type SwitchEnum string

const Switch = SwitchEnum("开关")

func (SwitchEnum) On() SwitchEnum {
	return "开启"
}
func (SwitchEnum) Off() SwitchEnum {
	return "关闭"
}
func (SwitchEnum) Enums() *goenum.Enums[SwitchEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(Switch.On()),
		goenum.NewEnum(Switch.Off()),
	)
}
