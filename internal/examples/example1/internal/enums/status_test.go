package enums_test

import (
	"testing"

	"github.com/yyle88/goenum/goenumgen"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateStatus(t *testing.T) {
	goenumgen.Generate(&goenumgen.Config[int]{
		Type:       "StatusEnum",
		Name:       "Status",
		BasicValue: 0,
		DelimValue: 0,
		Options: []*goenumgen.EnumOption[int]{
			{
				Name:        "OK",
				OptionValue: 200,
			},
			{
				Name:        "WA",
				OptionValue: 400,
			},
		},
		NamingMode: goenumgen.NamingMode.Single(),
		IsGenBasic: false,
		IsGenValid: false,
		IsGenCheck: false,
	}, runtestpath.SrcPath(t))
}
