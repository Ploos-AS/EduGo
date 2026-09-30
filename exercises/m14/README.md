# M14 — Architecture and larger Go programs

M14 covers package boundaries, `internal/`, dependency direction, consumer-owned interfaces, composition roots, dependency injection without frameworks, and configuration/startup design.

`taskapp` keeps its application service independent of transport and concrete storage. The `task` package owns the tiny interface it consumes, while `internal/memory` supplies one concrete adapter. Tests use a fake directly and require no network, database, globals, or DI container.
