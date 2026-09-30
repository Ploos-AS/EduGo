# 25. Fuzzing

Go integrates fuzzing with its test tool.

## Mission
Fuzz `ParsePort`. Arbitrary text must never cause a panic, and every accepted result must be in the range 1 through 65535.

Begin with useful seed cases and let the fuzzer explore additional inputs.
