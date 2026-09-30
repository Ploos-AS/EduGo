# 79. Pooling

`sync.Pool` kan redusere press fra midlertidige objekter, men er ikke en generell objektcache og innholdet kan forsvinne når som helst.

## Oppdrag
Benchmark en midlertidig buffer med og uten pool. Behold pooling bare dersom målingen viser en relevant gevinst.
