# 6. Error handling

Errors in Go are ordinary values. Functions commonly return a result together with an `error`.

Learn to create and return errors, check `err != nil`, use `errors.Is`, and wrap errors with `fmt.Errorf` and `%w`.

## Mission
Implement `ParsePort(string) (int, error)`. Accept ports 1 through 65535 and reject empty, non-numeric, and out-of-range input.
