# 57. TLS in practice

Use `httptest.NewTLSServer` for local HTTPS testing. Study trust and hostname verification without disabling verification as a shortcut.

## Mission
Serve HTTPS locally and use the test server's configured client to verify a request.
