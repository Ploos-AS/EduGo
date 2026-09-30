# 96. Composition root

Program startup is where concrete components are assembled.

## Mission
Let `main` load configuration, construct dependencies, build the application/service, and start the transport. Keep business logic out of `main`.
