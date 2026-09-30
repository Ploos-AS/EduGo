# 11. select and context

Use `select` to wait on multiple channel operations and `context.Context` for cancellation and deadlines.

## Mission
Create a worker producing periodic results that stops promptly when its context is cancelled.

Pass contexts as parameters and design goroutines so they have a clear termination path.
