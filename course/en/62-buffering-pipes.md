# 62. Buffering and pipes

Buffering can reduce system calls but changes latency and visibility of writes. Pipes connect producers and consumers as streams.

## Mission
Compare direct writes with `bufio.Writer`, then build an `io.Pipe` with concurrent producer and consumer.
