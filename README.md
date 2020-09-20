[![GoDoc](https://godoc.org/github.com/subbuv26/chanup?status.svg)](https://pkg.go.dev/github.com/subbuv26/chanup)
[![Go Report Card](https://goreportcard.com/badge/github.com/subbuv26/chanup)](https://goreportcard.com/report/github.com/subbuv26/chanup)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# ChanUp Channels
A Wrapper on top of go channels which supports one complex use case. 
A ChanUp channel buffer has a length of One.
ChanUp channel can be used to make producer and consumer of channel independent of each other.
The execution flow of no process gets blocked. 

## Functionality
- ChanUp channel never blocks producer or consumer.
- A producer Puts a value if channel is empty
- A producer Updates channel with new value if it is holding a stale value.
- A consumer gets a value if available.

## Installation
```
go get github.com/subbuv26/chanup
```

## Usage
### Example 1:
```go
package main

import (
	"fmt"
	"github.com/subbuv26/chanup"
)

type testType struct {
	a int
	s string
}

func main() {
	ch := chanup.GetChan()
	status := ch.Put(testType{
		a: 10,
		s: "Sample",
	})
	if status == chanup.FAILED {
		// Log
	}
	
	testValue := testType{
		a: 20,
		s: "Sample2",
	}
	status = ch.Update(testValue)

	if status != chanup.UPDATE {
		// Log
	}

	val := ch.Get()
	if val == nil {
		// Log
	}
	tv := val.(testType)
	fmt.Println(tv)
}
```

### Example 2:
```go
package main

import (
	"fmt"
	"github.com/subbuv26/chanup"
	"time"
)

type testType struct {
	a int
	s string
}

func producer(chp *chanup.ChanUp) {
	tv := testType{
		100,
		"sample",
	}
	for {
		time.Sleep(time.Second)
		chp.Update(tv)
		tv.a += 1
	}
}

func consumer(chc *chanup.ChanUp) {
	for {
		if val := chc.Get(); val != nil {
			tv := val.(testType)
			fmt.Println(tv)
			time.Sleep(time.Second * 5)
		}
	}
}

func main() {
	ch := chanup.GetChan()

	go producer(ch)

	consumer(ch)
}

```