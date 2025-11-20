// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: permission_test.go:11 -> enums_test.TestGeneratePermission
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"slices"
)

type PermissionEnum string

const Permission = PermissionEnum("權限")

func (PermissionEnum) E開啟() PermissionEnum {
	return "權限" + "-" + "開啟"
}
func (PermissionEnum) D關閉() PermissionEnum {
	return "權限" + "-" + "關閉"
}
func (PermissionEnum) Basic() PermissionEnum {
	return Permission
}
func (PermissionEnum) Enums() []PermissionEnum {
	return []PermissionEnum{
		Permission.E開啟(),
		Permission.D關閉(),
	}
}
func (value PermissionEnum) Valid() bool {
	return slices.Contains(Permission.Enums(), value)
}
func (value PermissionEnum) Check() bool {
	return value == Permission || slices.Contains(Permission.Enums(), value)
}
