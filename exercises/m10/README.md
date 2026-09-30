# M10 — Systems programming

M10 covers Go's portable systems-programming layer: streaming I/O, filesystems, buffering and pipes, subprocesses, signals, environment/configuration, and platform-specific boundaries.

`streamcopy` demonstrates composition through `io.Reader` and `io.Writer` plus an explicit input bound. The exercise intentionally avoids OS-specific APIs so the mandatory path remains portable.
