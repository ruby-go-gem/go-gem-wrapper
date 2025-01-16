# go-gem-wrapper
`go-gem-wrapper` is a wrapper for creating Ruby native extension in [Go](https://go.dev/)

[![GitHub Tag](https://img.shields.io/github/v/tag/ruby-go-gem/go-gem-wrapper)](https://github.com/ruby-go-gem/go-gem-wrapper/releases)
[![build](https://github.com/ruby-go-gem/go-gem-wrapper/actions/workflows/build.yml/badge.svg)](https://github.com/ruby-go-gem/go-gem-wrapper/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/ruby-go-gem/go-gem-wrapper/badge.svg)](https://coveralls.io/github/ruby-go-gem/go-gem-wrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/ruby-go-gem/go-gem-wrapper)](https://goreportcard.com/report/github.com/ruby-go-gem/go-gem-wrapper)
[![Go Reference](https://pkg.go.dev/badge/github.com/ruby-go-gem/go-gem-wrapper.svg)](https://pkg.go.dev/github.com/ruby-go-gem/go-gem-wrapper)

## Overview
| Directory        | Name                                                  | API Reference                                            |
|------------------|-------------------------------------------------------|----------------------------------------------------------|
| [/](/)           | `github.com/ruby-go-gem/go-gem-wrapper` (Go module)   | https://pkg.go.dev/github.com/ruby-go-gem/go-gem-wrapper |
| [/_gem/](/_gem/) | [go_gem](https://rubygems.org/gems/go_gem) (Ruby gem) | https://ruby-go-gem.github.io/go-gem-wrapper/            |

## Requirements
* Go 1.23+
* Ruby 3.3+

See [.github/workflows/matrix.json](.github/workflows/matrix.json) for details

## Getting started
At first, patch to make a gem into a Go gem right after `bundle gem`

See [_tools/patch_for_go_gem/](_tools/patch_for_go_gem/)

Please also add the following depending on the CI you are using.

### GitHub Actions
e.g.

```yml
- uses: actions/setup-go@v5
  with:
    go-version-file: ext/GEM_NAME/go.mod
```

## Implementing Ruby methods in Go
For example, consider the following Ruby method implemented in Go

```ruby
module Example
  def self.sum(a, b)
    a + b
  end
end
```

### 1. Implementing function in Go
```go
// ext/GEM_NAME/GEM_NAME.go

//export rb_example_sum
func rb_example_sum(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	aLong := ruby.NUM2LONG(ruby.VALUE(a))
	bLong := ruby.NUM2LONG(ruby.VALUE(b))

	sum := aLong + bLong

	return C.VALUE(ruby.LONG2NUM(sum))
}
```

### 2. Write C function definitions for Go functions
```go
// ext/GEM_NAME/GEM_NAME.go

/*
#include "example.h"

// TODO: Append this
VALUE rb_example_sum(VALUE self, VALUE a, VALUE b);
*/
import "C"
```

### 3. Call exported C functions with the Init function
```go
// ext/GEM_NAME/GEM_NAME.go

//export Init_example
func Init_example() {
	rb_mExample := ruby.RbDefineModule("Example")

	// TODO: Append this
	ruby.RbDefineSingletonMethod(rb_mExample, "sum", C.rb_example_sum, 2)
}
```

### More examples
See also

* [ruby/testdata/example/ext/example/example.go](ruby/testdata/example/ext/example/example.go)
* [ruby/testdata/example/ext/example/tests.go](ruby/testdata/example/ext/example/tests.go)

## Coverage
We provide auto-generated bindings for (almost) all CRuby functions available when including `ruby.h` :muscle:

See below for details.

* [ruby/enum_ruby_3_3_generated.go](ruby/enum_ruby_3_3_generated.go)
* [ruby/function_ruby_3_3_generated.go](ruby/function_ruby_3_3_generated.go)
* [ruby/type_ruby_3_3_generated.go](ruby/type_ruby_3_3_generated.go)
* [_tools/ruby_h_to_go/](_tools/ruby_h_to_go/)

## Specification
### Method name mapping from CRuby to Go
CRuby methods are mapped to Go methods based on the following rules

* No lowercase letters included (`/^[A-Z0-9_]+$/`)
  * No changes
  * e.g. `RB_NUM2UINT` (CRuby) -> `ruby.RB_NUM2UINT` (Go)
* Lowercase letters included
  * Converted to CamelCase
  * e.g. `rb_define_method` (CRuby) -> `ruby.RbDefineMethod` (Go)

### Limitation
Most of the methods defined in `ruby.h` are automatically generated and defined in [ruby/function_ruby_3_3_generated.go](ruby/function_ruby_3_3_generated.go).

However, some of the methods listed below are not supported.

1. deprecated or internal methods
    * See `function.exclude_name` in https://github.com/ruby-go-gem/ruby_header_parser/blob/main/config/default.yml
2. Methods with variable-length arguments
    * Because Go's variable-length arguments couldn't be passed directly to C.
    * However, it is possible to execute functions with variable length arguments in CRuby from Go with a hack like `RbRaise` in [ruby/ruby_internal_error.go](ruby/ruby_internal_error.go)

## Developing
### Build
Run `bundle exec rake build_all`.

See `bundle exec rake -T` for more tasks.

### See `godoc` in local
```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc
```

open http://localhost:6060/pkg/github.com/ruby-go-gem/go-gem-wrapper/ruby/

## Original idea
[Ruby meets Go - RubyKaigi 2015](https://rubykaigi.org/2015/presentations/mmasaki/)
