// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: action_test.go:11 -> enums_test.TestGenerateAction
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"slices"
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
func (ActionEnum) Basic() ActionEnum {
	return Action
}
func (ActionEnum) Enums() []ActionEnum {
	return []ActionEnum{
		Action.Start(),
		Action.Close(),
		Action.Pause(),
		Action.Reset(),
	}
}
func (value ActionEnum) Valid() bool {
	return slices.Contains(Action.Enums(), value)
}
func (value ActionEnum) Check() bool {
	return value == Action || slices.Contains(Action.Enums(), value)
}
