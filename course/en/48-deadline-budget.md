# 48. Deadlines and time budgets

A request crossing multiple services needs one overall time budget. Propagate `context.Context` through the call chain and preserve time for cleanup and response handling.

## Mission
Build local A -> B -> C calls and demonstrate cancellation propagating from A.
