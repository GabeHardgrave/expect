# expect
Runtime assertions in Go

## Installation
```sh
go get -u github.com/gabehardgrave/expect
```

## Usage

Add runtime assertions to your go code.
```go
var x = getStrictlyPositiveInt()
expect.True(x > 0, "%d wasn't positive!", x)
```
When run normally, `expect.True(...)` will `panic` if the invariant isn't met. Building with `-tags=release` will remove all `expect.True(...)` calls.