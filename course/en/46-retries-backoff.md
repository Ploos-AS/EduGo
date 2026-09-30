# 46. Retries and backoff

Retries can mask transient failures but can also amplify overload. Bound attempts, respect context/deadlines, and retry only when operation semantics permit it.

## Mission
Build a testable retry helper whose waiting mechanism can be injected so tests require no real sleeping.
