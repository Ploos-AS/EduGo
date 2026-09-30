# 48. Deadlines og tidsbudsjett

En request som går gjennom flere tjenester må ha et samlet tidsbudsjett.

Propager `context.Context` gjennom kallkjeden og reserver tid til opprydding og respons.

## Oppdrag
Lag A -> B -> C lokalt og vis hvordan cancellation fra A stopper underliggende arbeid.
