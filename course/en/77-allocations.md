# 77. Allocations and escape

Heap allocation is not automatically a problem, but unnecessary allocations can matter on hot paths. Study `-benchmem`, escape analysis, and API effects.

## Mission
Find one allocation in a benchmark, reduce it without changing results, and document before/after measurements.
