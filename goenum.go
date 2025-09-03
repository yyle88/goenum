// Package goenum: Go enumeration validation and type checking utilities
// Auto generates enum validation functions with type safety guarantees
// Supports generic enum types and provides compile-time enum verification
// Enables efficient enum validation with zero runtime allocations
//
// goenum: Go 枚举验证和类型检查工具包
// 自动生成具有类型安全保证的枚举验证函数
// 支持泛型枚举类型，提供编译时枚举验证
// 实现零运行时分配的高效枚举验证
package goenum

import "github.com/yyle88/goenum/internal/constraint"

// Valid checks if the given enum value exists in the enum set
// Uses compile-time type safety with generic constraints
// Returns true if value is found in the enum collection
//
// Valid 检查给定的枚举值是否存在于枚举集合中
// 使用泛型约束实现编译时类型安全
// 如果值在枚举集合中找到则返回 true
func Valid[E constraint.EnumType[E]](value E) bool {
	// Iterate through all enum values to find match
	// 遍历所有枚举值以查找匹配项
	for _, elem := range value.Enums() {
		if elem == value {
			return true
		}
	}
	return false
}

// Check validates enum value with basic value fallback support
// Accepts both basic enum values and valid enum options
// Combines basic value check with comprehensive enum validation
//
// Check 验证枚举值，支持基本值回退
// 接受基本枚举值和有效的枚举选项
// 结合基本值检查和全面的枚举验证
func Check[E constraint.EnumItem[E]](value E) bool {
	return value == value.Basic() || Valid(value)
}
