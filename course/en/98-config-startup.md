# 98. Configuration and startup

Separate configuration parsing, validation, and runtime use. Invalid critical configuration should fail startup clearly.

## Mission
Create an immutable configuration value from environment/flags, validate it once, and pass it explicitly to components that need it.
