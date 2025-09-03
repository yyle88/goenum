package enums

import (
	"slices"
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
func (ConnectionEnum) Basic() ConnectionEnum {
	return Connection
}
func (ConnectionEnum) Enums() []ConnectionEnum {
	return []ConnectionEnum{
		Connection.C接続(),
		Connection.D切断(),
		Connection.W待機(),
	}
}
func (value ConnectionEnum) Valid() bool {
	return slices.Contains(Connection.Enums(), value)
}
func (value ConnectionEnum) Check() bool {
	return value == Connection || slices.Contains(Connection.Enums(), value)
}
