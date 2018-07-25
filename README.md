# kpass

## Overview
kpass is validation tool.  
each column, record, etc.. is verified that expected value.  

### Usage
install
```
go get -u github.com/midorigreen/kpass
```

build
```sh
go build
```

run
```sh
(stdin) > ./kpass
```

### Extention (Not Using)
Adding new validation is 3 steps.

1. Implement Gater interface struct
```go
// gate/gate.go

type Gater interface {
	Examine(c Condition, e string, r chan Result)
}
```

ex)
```go
type Minimumer struct{}

func (m Minimumer) Examine(c Condition, e string, r chan Result) {
	// examine
}
```

2. Add implemented struct to slice
```go
// check.go

var gaters = []gate.Gater{
	gate.Minimumer{},
}
```

(making)
3. Create `gate.Condition` struct
4. Execute
