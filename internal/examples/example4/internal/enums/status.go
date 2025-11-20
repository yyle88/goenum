// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: status_test.go:11 -> enums_test.TestGenerateStatus
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"slices"
)

type StatusEnum string

const Status = StatusEnum("状态")

func (StatusEnum) S成功() StatusEnum {
	return "状态" + "-" + "成功"
}
func (StatusEnum) F失败() StatusEnum {
	return "状态" + "-" + "失败"
}
func (StatusEnum) P等待() StatusEnum {
	return "状态" + "-" + "等待"
}
func (StatusEnum) Basic() StatusEnum {
	return Status
}
func (StatusEnum) Enums() []StatusEnum {
	return []StatusEnum{
		Status.S成功(),
		Status.F失败(),
		Status.P等待(),
	}
}
func (value StatusEnum) Valid() bool {
	return slices.Contains(Status.Enums(), value)
}
func (value StatusEnum) Check() bool {
	return value == Status || slices.Contains(Status.Enums(), value)
}
