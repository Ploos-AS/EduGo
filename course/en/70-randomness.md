# 70. Secure randomness

Use `crypto/rand` when values must be unpredictable. `math/rand` is useful for simulation and tests but is not a substitute for cryptographic randomness.

## Mission
Generate a token from a fixed number of random bytes and encode it using URL-safe base64.
