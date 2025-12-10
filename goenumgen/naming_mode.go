// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: naming_mode_test.go:11 -> goenumgen_test.TestGenerateAction
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package goenumgen

import (
	"github.com/yyle88/goenum"
)

type NamingModeEnum string

const NamingMode = NamingModeEnum("Naming")

func (NamingModeEnum) Prefix() NamingModeEnum {
	return "Naming" + "-" + "Prefix"
}
func (NamingModeEnum) Suffix() NamingModeEnum {
	return "Naming" + "-" + "Suffix"
}
func (NamingModeEnum) Middle() NamingModeEnum {
	return "Naming" + "-" + "Middle"
}
func (NamingModeEnum) Single() NamingModeEnum {
	return "Naming" + "-" + "Single"
}
func (NamingModeEnum) Enums() *goenum.Enums[NamingModeEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(NamingMode.Prefix()),
		goenum.NewEnum(NamingMode.Suffix()),
		goenum.NewEnum(NamingMode.Middle()),
		goenum.NewEnum(NamingMode.Single()),
	)
}
