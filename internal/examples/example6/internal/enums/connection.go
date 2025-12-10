// Code generated using goenumgen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/goenum
// Generated from: connection_test.go:11 -> enums_test.TestGenerateConnection
// ========== GOENUM:DO-NOT-EDIT-SECTION:END ==========

package enums

import (
	"github.com/yyle88/goenum"
)

type ConnectionEnum string

const Connection = ConnectionEnum("接続")

func (ConnectionEnum) C接続() ConnectionEnum {
	return "接続" + "" + "接続"
}
func (ConnectionEnum) D切断() ConnectionEnum {
	return "接続" + "" + "切断"
}
func (ConnectionEnum) W待機() ConnectionEnum {
	return "接続" + "" + "待機"
}
func (ConnectionEnum) Enums() *goenum.Enums[ConnectionEnum] {
	return goenum.NewEnums(
		goenum.NewEnum(Connection.C接続()),
		goenum.NewEnum(Connection.D切断()),
		goenum.NewEnum(Connection.W待機()),
	)
}
