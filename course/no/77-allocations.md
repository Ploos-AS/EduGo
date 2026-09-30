# 77. Allocations og escape

Heap-allokering er ikke automatisk et problem, men unødvendige allocations kan bli kostbare i varme kodebaner.

Studer `-benchmem`, escape analysis og hvordan API-design påvirker allocations.

## Oppdrag
Finn en allocation i en benchmark, reduser den uten å endre resultatet, og dokumenter før/etter-målingen.
