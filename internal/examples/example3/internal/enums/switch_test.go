package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateSwitch(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[string]{
		Type:       "SwitchEnum",
		Name:       "Switch",
		BasicValue: "开关",
		DelimValue: "-",
		Options: []*goenumgen.EnumOption[string]{
			{
				Name:        "On",
				OptionValue: "开启",
			},
			{
				Name:        "Off",
				OptionValue: "关闭",
			},
		},
		NamingMode: goenumgen.NamingMode.Single(),
	}, runtestpath.SrcPath(t))
}
