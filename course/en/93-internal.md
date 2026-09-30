# 93. internal and public API

Go enforces import boundaries for directories beneath `internal/`. Use this to keep implementation details private while stable entry points remain small.

## Mission
Move an implementation detail behind `internal/` and document which surface is actually a public contract.
