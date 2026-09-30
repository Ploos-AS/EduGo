# 19. Logging and configuration

Use structured logging with `log/slog`. Read configuration from flags and environment variables and validate it before starting work.

## Mission
Extend the M3 service with `PORT`, `LOG_LEVEL`, and command-line overrides.

Never log secrets; distinguish operational logs from user-facing CLI output.
