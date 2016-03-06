# prob
Probability experiments in Go

[![GoDoc](https://godoc.org/github.com/brnstz/prob?status.svg)](https://godoc.org/github.com/brnstz/prob)
[![Build Status](https://travis-ci.org/brnstz/prob.svg?branch=master)](https://travis-ci.org/brnstz/prob.svg?branch=master)

## Requirements

* Go 1.6+

## Example

```go
package main

import (
    "fmt"
    "math/big"

    "github.com/brnstz/prob"
)

func main() {
    // There are 100 restaurants in a city. A restaurant reviewer wants to
    // eat dinner at a different restaurant every night in a week. How many
    // different possible dining schedules could he have?
    var (
        restaurants int64 = 100
        days        int64 = 7
    )

    schedules := prob.NoReplaceOrdered(big.NewInt(restaurants), big.NewInt(days))

    fmt.Printf(
        "There are %s different ways to schedule dinner.",
        schedules.Text(10),
    )

}
```
