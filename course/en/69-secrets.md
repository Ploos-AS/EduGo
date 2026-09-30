# 69. Secrets and sensitive configuration

Do not hard-code, commit, or log secrets. Read them at the process boundary and avoid spreading them through unnecessary data structures.

## Mission
Load a required secret from the environment and return a clear error when absent without including its value.
