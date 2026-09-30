# EduGo

**Learn Go from zero to practical systems and network programming.**

EduGo is a bilingual (Norwegian/English) course from Ploos AS. It is designed for learners who may be completely new to Go, while still going far enough to build useful command-line tools, network services, concurrent programs, and production-style applications.

## Goals

By the end of the course, the learner should be able to:

- read and write idiomatic Go;
- understand variables, types, control flow, functions, structs, interfaces, slices, maps, pointers, errors, packages, and modules;
- use goroutines, channels, contexts, and synchronization safely;
- build command-line applications;
- work with files, JSON, HTTP, TCP, and basic network protocols;
- test, benchmark, format, vet, and document Go code;
- understand the Go toolchain and module ecosystem;
- build small real-world services and utilities;
- understand where Go differs from C and Python.

## Course structure

The canonical course sources live under `course/`:

- `course/no/` — Norwegian primary edition
- `course/en/` — English edition
- `examples/` — runnable examples
- `exercises/` — student exercises
- `solutions/` — reference solutions
- `projects/` — larger milestone projects

The course starts with fundamentals and develops toward practical systems, networking, concurrency, and service programming.

## Student environment

Students must not depend on internal Ploos infrastructure. EduGo therefore includes a self-contained student OCI image definition under `student-oci/`. A local Go installation is also supported.

## Suggested learning path

1. First Go program and toolchain
2. Values, variables, constants, and types
3. Control flow
4. Functions
5. Arrays, slices, and maps
6. Structs, methods, and interfaces
7. Pointers and memory model basics
8. Errors and defensive programming
9. Packages and modules
10. Files and structured data
11. Testing, benchmarks, and tooling
12. Goroutines and channels
13. Context, cancellation, and synchronization
14. Command-line applications
15. HTTP clients and servers
16. TCP and network programming
17. Persistence and configuration
18. Observability and graceful shutdown
19. Profiling and performance
20. Capstone project

## Publishing

EduGo is intended for single-source publishing to web/HTML, EPUB, Kindle-compatible output, and PDF using the Ploos publishing workflow.

## License

Course material and documentation are licensed under **CC BY 4.0** unless a file explicitly states otherwise.

Copyright © 2026 Ploos AS.
