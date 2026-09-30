# 66. Platform-specific code

Go can select files by GOOS/GOARCH and build constraints.

## Mission
Create a package with one portable API and separate implementations only where the operating system requires them. Keep platform differences at the system boundary.
