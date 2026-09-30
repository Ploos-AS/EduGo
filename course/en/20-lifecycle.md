# 20. Process lifecycle and graceful shutdown

Handle OS signals with `signal.NotifyContext`, stop accepting new work, allow active work to finish within a deadline, and release resources.

## Mission
Extend miniservice so SIGINT/SIGTERM triggers graceful shutdown with an explicit timeout.
