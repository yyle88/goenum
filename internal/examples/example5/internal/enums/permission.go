// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: permission_test.go:11 -> enums_test.TestGeneratePermission
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"github.com/yyle88/goenum"
)

type PermissionEnum string

const Permission = PermissionEnum("權限")

func (PermissionEnum) E開啟() PermissionEnum {
	return "權限" + "-" + "開啟"
}
func (PermissionEnum) D關閉() PermissionEnum {
	return "權限" + "-" + "關閉"
}
func (PermissionEnum) Enums() *goenum.Enums[PermissionEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(Permission.E開啟()),
		goenum.NewEnum(Permission.D關閉()),
	)
}
