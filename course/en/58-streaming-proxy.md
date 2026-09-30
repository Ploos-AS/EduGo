# 58. Streaming and proxying

Streaming moves data progressively instead of buffering everything first. Proxies must handle cancellation, backpressure, headers, and errors on both sides.

## Mission
Build a local streaming handler with `io.Copy` and test that client cancellation stops work.
