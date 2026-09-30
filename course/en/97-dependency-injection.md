# 97. Dependency injection without a framework

Constructors and ordinary Go values are often sufficient for dependency injection.

## Mission
Replace a real dependency with a small fake in a test without global state or a DI container. Create interfaces only when they provide a useful boundary.
