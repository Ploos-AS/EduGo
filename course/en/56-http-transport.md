# 56. HTTP transport internals

Study connection reuse, keep-alive, request bodies, response bodies, and why clients normally reuse `http.Client` and its Transport rather than creating a new transport per request.

## Mission
Use `httptest.Server` and instrumentation to inspect multiple requests to one local server.
