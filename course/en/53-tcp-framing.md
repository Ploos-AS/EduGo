# 53. TCP in depth: framing

TCP provides an ordered byte stream, not application messages. Study delimiter, length-prefix, and fixed-size framing with explicit size bounds.

## Mission
Implement a line-based PING/ECHO protocol over `net.Conn` and test it with `net.Pipe`.
