# 45. Protocol design

A protocol needs explicit message boundaries, versioning, size limits, and defined errors.

## Mission
Design a small text protocol with `PING` and `ECHO`. Define a maximum line size and rejection behavior. Remember that TCP is a byte stream: one `Read` is not necessarily one message.
