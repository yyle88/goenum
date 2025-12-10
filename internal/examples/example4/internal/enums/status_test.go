package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateStatus(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "StatusEnum",
		Name:       "Status",
		BasicValue: "状态",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{
				Name:        "S成功",
				OptionValue: "成功",
			},
			{
				Name:        "F失败",
				OptionValue: "失败",
			},
			{
				Name:        "P等待",
				OptionValue: "等待",
			},
		},
		NamingMode: goenumgen.NamingMode.Suffix(),
	}, runtestpath.SrcPath(t))
}
