# 22. Dependencies and testable design

Pass dependencies explicitly through constructors and small interfaces. Keep application wiring in `main` separate from domain behavior.

## Mission
Give miniservice a logger and a small storage abstraction without mutable global state.

Define interfaces where consumers need abstraction rather than wrapping every concrete type automatically.
