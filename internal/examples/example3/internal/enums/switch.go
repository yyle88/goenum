package enums

import (
	"slices"
)

type SwitchEnum string

const Switch = SwitchEnum("开关")

func (SwitchEnum) On() SwitchEnum {
	return "开启"
}
func (SwitchEnum) Off() SwitchEnum {
	return "关闭"
}
func (SwitchEnum) Basic() SwitchEnum {
	return Switch
}
func (SwitchEnum) Enums() []SwitchEnum {
	return []SwitchEnum{
		Switch.On(),
		Switch.Off(),
	}
}
func (value SwitchEnum) Valid() bool {
	return slices.Contains(Switch.Enums(), value)
}
func (value SwitchEnum) Check() bool {
	return value == Switch || slices.Contains(Switch.Enums(), value)
}
