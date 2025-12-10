// Package goenum: Go enumeration code generation and type safe enum utilities
// Auto generates enum types with compile-time type safety guarantees
// Integrates with github.com/yyle88/enum package for collection management
// Enables efficient enum validation with map-based lookup
//
// goenum: Go 枚举代码生成和类型安全枚举工具包
// 自动生成具有编译时类型安全保证的枚举类型
// 与 github.com/yyle88/enum 包集成进行集合管理
// 实现基于 map 查找的高效枚举验证
package goenum

import "github.com/yyle88/enum"

// Enum is a type alias for enum.Enum with MetaNone metadata
// Simplifies the generic signature from two type params to one
//
// Enum 是 enum.Enum 的类型别名，使用 MetaNone 元数据
// 将泛型签名从两个类型参数简化为一个
type Enum[enumCode comparable] = enum.Enum[enumCode, *enum.MetaNone]

// NewEnum creates a new Enum instance with Go native enum
// Wraps enum.NewEnum for convenience
//
// 创建新的 Enum 实例，绑定 Go 原生枚举
// 封装 enum.NewEnum 以便使用
func NewEnum[C comparable](code C) *Enum[C] {
	return enum.NewEnum(code)
}

// Enums is a type alias for enum.Enums with MetaNone metadata
// Simplifies the generic signature from two type params to one
//
// Enums 是 enum.Enums 的类型别名，使用 MetaNone 元数据
// 将泛型签名从两个类型参数简化为一个
type Enums[enumCode comparable] = enum.Enums[enumCode, *enum.MetaNone]

// NewEnums creates a new Enums collection from the given Enum instances
// Wraps enum.NewEnums for convenience
//
// 从给定的 Enum 实例创建新的 Enums 集合
// 封装 enum.NewEnums 以便使用
func NewEnums[C comparable](params ...*Enum[C]) *Enums[C] {
	return enum.NewEnums(params...)
}
