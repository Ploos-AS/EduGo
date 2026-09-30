# 94. Dependency direction

Architecture is more about dependency direction than the number of layers. Domain logic should not know about HTTP, SQL, or command-line flags unless it needs them.

## Mission
Draw imports in a small service and move one technical dependency out of core logic.
