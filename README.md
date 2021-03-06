go-tskv
=======
Go library for Time Series Key Value storage.

[![Build Status](https://travis-ci.org/brimstone/go-tskv.svg?branch=master)](https://travis-ci.org/brimstone/go-tskv) [![Coverage Status](https://coveralls.io/repos/github/brimstone/go-tskv/badge.svg?branch=master)](https://coveralls.io/github/brimstone/go-tskv?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/brimstone/go-tskv)](https://goreportcard.com/report/github.com/brimstone/go-tskv)

Usage
-----

```
package main

import (
	"fmt"

	"github.com/brimstone/go-tskv"
)

func main() {
	weight, err := tskv.New(tskv.Config{
		Duration:  time.Hour,
		Frequency: time.Minute,
	})
	if err != nil {
		panic(err)
	}
	weight.InsertNow(100)
	fmt.Println("Last weight:", weight.Last().Int())
}
```

