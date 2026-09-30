# 0. Introduction

Welcome to EduGo.

The goal of this course is to take you from being new to Go to building useful command-line programs, network tools, and services on your own.

## What you need

You can use either:

- a local Go installation, or
- the student environment in `student-oci/`.

The course is designed so that you do not need any internal Ploos infrastructure.

## How to work through the course

Each chapter follows roughly the same rhythm:

1. Read the concept.
2. Run the examples.
3. Change them.
4. Solve the exercises.
5. Build a small part of a larger project.

## Your first program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it with:

```sh
go run .
```

The next chapter introduces the Go toolchain, `go mod`, `go fmt`, `go test`, and the structure of a Go project.
