# 27. Go memory model

Learn practical happens-before reasoning and the synchronization guarantees provided by channels, mutexes, and atomics.

## Mission
Start with unsynchronized shared state. Explain why assuming another goroutine will observe a write is invalid, then repair the design in two ways.

Treat the documented memory model as the contract, not accidental scheduler observations.
