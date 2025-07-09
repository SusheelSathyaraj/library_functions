## `atoi` — Convert string to integer

This is a manual reimplementation of Go's `strconv.Atoi()` function.

### Features
- Handles leading/trailing spaces
- Parses optional `+` or `-` signs
- Stops parsing at the first non-digit
- Detects and clamps overflow

### Example

```go
fmt.Println(myAtoi("   -42"))        // Output: -42
fmt.Println(myAtoi("4193 with text")) // Output: 4193
