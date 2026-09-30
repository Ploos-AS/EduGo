# 104. Capstone M4 — Verification

Run unit, integration, race, and negative tests. Add fuzzing where parsers or input boundaries benefit from it.

## Acceptance
`go test ./...`, `go test -race ./...`, and formatting checks pass. Benchmark only a path backed by a concrete performance hypothesis.
