# M11 — Defensive security engineering

M11 covers secure defaults, secret handling, cryptographic randomness, hashing versus MACs, TLS configuration, dependency hygiene, and security-focused testing.

`authtoken` is deliberately small and uses only Go's standard cryptographic library. It generates random opaque payloads and authenticates them with HMAC-SHA-256; tests verify normal operation, tampering rejection, and key validation.
