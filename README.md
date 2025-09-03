[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goenum/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goenum/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goenum)](https://pkg.go.dev/github.com/yyle88/goenum)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goenum/main.svg)](https://coveralls.io/github/yyle88/goenum?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goenum.svg)](https://github.com/yyle88/goenum/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goenum)](https://goreportcard.com/report/github.com/yyle88/goenum)

# goenum

Go enumeration generation and management toolkit with type safety and flexible naming patterns.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Key Features

🎯 **Smart Enum Generation**: Auto generates type-safe enum code with customizable naming patterns  
⚡ **Multiple Naming Modes**: Supports prefix, suffix, middle, and single naming strategies  
🔄 **Type Safety**: Compile-time enum validation with generic constraints  
🌍 **Flexible Types**: Works with any comparable type (int, string, custom types)  
📋 **Validation Functions**: Auto-generates Valid() and Check() methods for runtime validation

## Installation

```bash
go get github.com/yyle88/goenum
```

## Quick Start

### 1. Define Your Enum Configuration

```go
package main

import (
    "github.com/yyle88/goenum/goenumgen"
)

func main() {
    // Configure enum generation
    config := &goenumgen.Config[string]{
        Type:       "StatusEnum",
        Name:       "Status", 
        BasicValue: "Status",
        DelimValue: "-",
        NamingMode: goenumgen.NamingMode.Suffix(), // "Status-OK", "Status-Error"
        IsGenValid: true,
        IsGenCheck: true,
        Options: []*goenumgen.EnumOption[string]{
            {Name: "OK", OptionValue: "OK"},
            {Name: "Error", OptionValue: "Error"}, 
            {Name: "Pending", OptionValue: "Pending"},
        },
    }
    
    // Generate enum code
    goenumgen.Generate(config, "internal/enums/status.go")
}
```

### 2. Generated Enum Code

The above configuration generates:

```go
package enums

import "slices"

type StatusEnum string

const Status = StatusEnum("Status")

func (StatusEnum) OK() StatusEnum {
    return "Status" + "-" + "OK"
}

func (StatusEnum) Error() StatusEnum {
    return "Status" + "-" + "Error"
}

func (StatusEnum) Pending() StatusEnum {
    return "Status" + "-" + "Pending"
}

func (StatusEnum) Enums() []StatusEnum {
    return []StatusEnum{
        Status.OK(),
        Status.Error(), 
        Status.Pending(),
    }
}

func (value StatusEnum) Valid() bool {
    return slices.Contains(Status.Enums(), value)
}

func (value StatusEnum) Check() bool {
    return value == Status || slices.Contains(Status.Enums(), value)
}
```

### 3. Use Your Enums

```go
package main

import (
    "fmt"
    "your-project/internal/enums"
    "github.com/yyle88/goenum"
)

func main() {
    // Create enum values
    status := enums.Status.OK()
    fmt.Println(status) // Output: Status-OK
    
    // Validate enum values
    if goenum.Valid(status) {
        fmt.Println("Valid enum value")
    }
    
    // Check with basic value support
    if goenum.Check(enums.Status) {
        fmt.Println("Basic value is valid")
    }
    
    // Get all enum values
    allStatuses := enums.Status.Enums()
    for _, s := range allStatuses {
        fmt.Printf("Status: %s, Valid: %t\n", s, s.Valid())
    }
}
```

## Naming Modes

### Prefix Mode
Pattern: `option + delimiter + basic`
```go
NamingMode: goenumgen.NamingMode.Prefix()
// Result: "OK-Status", "Error-Status"
```

### Suffix Mode  
Pattern: `basic + delimiter + option`
```go
NamingMode: goenumgen.NamingMode.Suffix()
// Result: "Status-OK", "Status-Error"
```

### Middle Mode
Pattern: `basic + option + delimiter`
```go
NamingMode: goenumgen.NamingMode.Middle()
// Result: "StatusOK-", "StatusError-"
```

### Single Mode
Pattern: `option`
```go
NamingMode: goenumgen.NamingMode.Single()
// Result: "OK", "Error"
```

## Advanced Examples

### HTTP Status Codes

```go
config := &goenumgen.Config[int]{
    Type:       "HTTPStatusEnum",
    Name:       "HTTPStatus",
    BasicValue: 0,
    DelimValue: 0, // Not used for integers
    NamingMode: goenumgen.NamingMode.Single(),
    IsGenValid: true,
    IsGenCheck: true,
    Options: []*goenumgen.EnumOption[int]{
        {Name: "OK", OptionValue: 200},
        {Name: "NotFound", OptionValue: 404},
        {Name: "InternalError", OptionValue: 500},
    },
}
```

### Database Connection States

```go
config := &goenumgen.Config[string]{
    Type:       "ConnStateEnum", 
    Name:       "ConnState",
    BasicValue: "conn",
    DelimValue: ".",
    NamingMode: goenumgen.NamingMode.Prefix(),
    IsGenBasic: true,
    IsGenValid: true,
    Options: []*goenumgen.EnumOption[string]{
        {Name: "Connected", OptionValue: "active"},
        {Name: "Disconnected", OptionValue: "inactive"},
        {Name: "Connecting", OptionValue: "pending"},
    },
}
// Generates: "active.conn", "inactive.conn", "pending.conn"
```

## Configuration Options

| Field | Type | Description |
|-------|------|-------------|
| `Type` | `string` | Generated enum type name |
| `Name` | `string` | Base constant name |  
| `BasicValue` | `T` | Basic value for the enum |
| `DelimValue` | `T` | Delimiter for compound names |
| `Options` | `[]*EnumOption[T]` | Enum option definitions |
| `NamingMode` | `NamingModeEnum` | Naming pattern strategy |
| `IsGenBasic` | `bool` | Generate `Basic()` method |
| `IsGenValid` | `bool` | Generate `Valid()` method | 
| `IsGenCheck` | `bool` | Generate `Check()` method |

## Validation Functions

### `goenum.Valid()`
Checks if a value exists in the enum set:
```go
if goenum.Valid(status) {
    // Value is one of the defined enum options
}
```

### `goenum.Check()`  
Validates with basic value fallback:
```go
if goenum.Check(status) {
    // Value is either basic value or valid enum option
}
```

## Project Structure

```
goenum/
├── goenum.go              # Main validation functions
├── goenumgen/             # Code generation package
│   ├── generate.go        # Generation engine
│   └── naming_mode.go     # Naming pattern definitions
├── internal/
│   ├── constraint/        # Generic type constraints
│   ├── utils/             # Utility functions
│   └── examples/          # Usage examples
│       ├── example1/      # Basic int enum
│       ├── example2/      # String enum with validation
│       ├── example3/      # Switch pattern enum
│       └── example4/      # Complex naming patterns
└── README.md
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-08-28 08:33:43.829511 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/goenum.svg?variant=adaptive)](https://starchart.cc/yyle88/goenum)
