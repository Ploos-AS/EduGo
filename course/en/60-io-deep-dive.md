# 60. io.Reader and io.Writer deep dive

`io.Reader` and `io.Writer` connect files, networks, compression, hashing, and processes without coupling components.

## Mission
Build a pipeline with `io.Copy`, `io.LimitReader`, and `io.TeeReader`. Explain where bytes move and where limits are enforced.
