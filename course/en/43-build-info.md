# 43. Build metadata and reproducibility

Programs should be able to report what they were built from. Explore `runtime/debug.ReadBuildInfo`, linker-injected values, and VCS information.

## Mission
Build a `version` command that reports useful build information without contaminating core logic.
