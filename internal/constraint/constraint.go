// Package constraint: Generic type constraints for enum validation and type safety
// Defines interface contracts for enum types with compile-time guarantees
// Supports both basic enum validation and extended enum item functionality
// Enables type-safe enum operations with zero runtime overhead
//
// constraint: 枚举验证和类型安全的泛型类型约束
// 定义枚举类型的接口合约，具有编译时保证
// 支持基本枚举验证和扩展枚举项功能
// 实现零运行时开销的类型安全枚举操作
package constraint

// EnumType defines the contract for basic enum validation
// Requires comparable type and method to get all enum values
// Used with Valid() function for runtime enum validation
//
// EnumType 定义基本枚举验证的合约
// 要求可比较类型和获取所有枚举值的方法
// 与 Valid() 函数结合用于运行时枚举验证
type EnumType[E comparable] interface {
	comparable
	Enums() []E
}

// EnumItem extends EnumType with basic value support
// Provides fallback to basic enum value for validation
// Used with Check() function for comprehensive enum validation
//
// EnumItem 扩展 EnumType，支持基本值
// 为验证提供基本枚举值回退
// 与 Check() 函数结合用于全面的枚举验证
type EnumItem[E comparable] interface {
	comparable
	Basic() E
	Enums() []E
}
