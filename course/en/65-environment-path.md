# 65. Environment, PATH, and configuration

Treat environment variables as process input rather than global application state.

## Mission
Load configuration into a struct at startup, validate it, and pass it explicitly. Test with `t.Setenv` without depending on the user's environment.
