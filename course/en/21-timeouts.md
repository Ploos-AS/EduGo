# 21. Timeouts, deadlines and bounds

Network software must not wait forever.

Configure explicit client/server timeouts, use context deadlines for cancellable work, and bound externally controlled input.

## Mission
Add HTTP timeouts and tests that simulate a slow handler.
