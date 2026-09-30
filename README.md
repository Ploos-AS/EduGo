# EduGo

**Learn Go from zero to production-ready systems and network programming.**

EduGo is a bilingual Norwegian/English course from Ploos AS. Norwegian is the primary edition and English is maintained in parallel. The course starts at zero and progresses through practical systems, networking, concurrency, service engineering, security, performance, operations, architecture, and a capstone service.

## Learning path

The numbered course currently runs through chapter 105 and is organized as milestones:

- **M0–M1:** orientation and Go fundamentals
- **M2:** concurrency
- **M3:** practical Go — CLI, files, HTTP, TCP/UDP, logging/config
- **M4:** production Go — lifecycle, timeouts, observability, profiling, fuzzing, distribution
- **M5:** language/runtime depth — memory model, stack/heap, scheduler/GC, generics, reflection
- **M6:** service engineering — persistence, TLS, auth, middleware, metrics, integration tests
- **M7:** API and release engineering
- **M8:** protocols and distributed-system fundamentals
- **M9:** network engineering
- **M10:** systems programming
- **M11:** defensive security engineering
- **M12:** performance engineering
- **M13:** observability and operations
- **M14:** architecture and larger Go programs
- **M15:** capstone Job Service and release

## Repository structure

- `course/no/` — Norwegian primary course
- `course/en/` — English parallel course
- `examples/` — runnable examples
- `exercises/` — focused student exercises
- `solutions/` — reference solutions where useful
- `projects/` — larger integrated projects and capstone
- `student-oci/` — self-contained student environment

Students do not need internal Ploos infrastructure. A local Go installation is supported, while the student OCI provides a reproducible Debian-based environment.

## Quality gates

CI verifies Go formatting, runs `go vet ./...`, tests every Go module, repeats tests with the race detector, and builds the student OCI image. Exercises should remain deterministic and avoid unnecessary external services.

## Publishing

The course is designed for single-source publishing to web/HTML, EPUB, Kindle-compatible output, and PDF through the Ploos publishing workflow.

## License

Course material and documentation are licensed under **CC BY 4.0** unless a file explicitly states otherwise.

Copyright © 2026 Ploos AS.
