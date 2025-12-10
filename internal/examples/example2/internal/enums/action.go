// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: action_test.go:11 -> enums_test.TestGenerateAction
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"github.com/yyle88/goenum"
)

type ActionEnum string

const Action = ActionEnum("Action")

func (ActionEnum) Start() ActionEnum {
	return "Action" + "-" + "Start"
}
func (ActionEnum) Close() ActionEnum {
	return "Action" + "-" + "Close"
}
func (ActionEnum) Pause() ActionEnum {
	return "Action" + "-" + "Pause"
}
func (ActionEnum) Reset() ActionEnum {
	return "Action" + "-" + "Reset"
}
func (ActionEnum) Enums() *goenum.Enums[ActionEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(Action.Start()),
		goenum.NewEnum(Action.Close()),
		goenum.NewEnum(Action.Pause()),
		goenum.NewEnum(Action.Reset()),
	)
}
