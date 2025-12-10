package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGeneratePermission(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "PermissionEnum",
		Name:       "Permission",
		BasicValue: "權限",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{Name: "E開啟", OptionValue: "開啟"},
			{Name: "D關閉", OptionValue: "關閉"},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
	}, runtestpath.SrcPath(t))
}
