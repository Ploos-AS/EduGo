# 40. Error modeling

Errors are part of an API. Use wrapping with `%w`, `errors.Is`, and `errors.As` when callers need to understand error categories. Separate expected domain errors from unexpected internal failures.

## Mission
Give noteservice machine-understandable domain errors without leaking internals over HTTP.
