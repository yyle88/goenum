// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: status_test.go:11 -> enums_test.TestGenerateStatus
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

type StatusEnum int

const Status = StatusEnum(0)

func (StatusEnum) OK() StatusEnum {
	return 200
}
func (StatusEnum) WA() StatusEnum {
	return 400
}
func (StatusEnum) Enums() []StatusEnum {
	return []StatusEnum{
		Status.OK(),
		Status.WA(),
	}
}
