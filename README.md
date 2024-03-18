# Good Enough IDs

This is a package that generates short IDs that are "unique" enough to be used instead of UUIDs, which are ugly.

Example: `ab-ce093c1`

## Features

- :sparkles: Optional prefixes: IDs will be of the form `prefix-*`
- :sparkles: Safe to use with goroutines: New IDs require data from a counter which has a lock
- :sparkles: Safe to use with multiple machines: Can set a unique "Machine ID" for each instance
- :sparkles: Optional custom epoch: Can override the default epoch of January 1st, 1970

## Install

```sh
go install github.com/jxlxx/geid@latest
```

## Getting Started

Epoch is optional. The default epoch is January 1st, 1970. Setting the epoch to a later date means short IDs.

MachineID is optional. The default machine ID is "1". This is useful if you have many instances of identical
ID generators running at the same time.

```go
package main

import (
	"fmt"
	"time"

	"github.com/jxlxx/geid"
)

type Cat struct{}

func (c Cat) Prefix() string {
	return "cat-"
}

func main() {
	c := Cat{}
	g := geid.New(c)

	t := time.Now()
	geid.SetCustomEpoch(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC))

	for i := 0; i < 100; i++ {
		fmt.Println(g.NewID())
	}
}
```

## Examples

```sh
cat-ab-ce093c1
cat-ac-ce093c1
cat-ad-ce093c1
```

## Design

The implementation is loosely based on the following article: [Creating User-Facing, Short Unique IDs: What are the options? - Hwee Lin Yeo, Alexandra](https://medium.com/teamocard/creating-user-facing-short-unique-id-ids-what-are-the-options-464a19283d98)

> Twitter Snowflake, the open-source version of which is unfortunately archived, is an internal service used by Twitter for generating 64-bit unique IDs at a high scale. The IDs are made up of the components:
> - Epoch timestamp in millisecond precision — 41 bits (gives us 69 years with a custom epoch)
> - Machine id — 10 bits (thus allowing uniqueness even when we scale the short ID generator service over different nodes)
> - Sequence number — 12 bits (A local counter per machine that rolls over every 4096)
> - The extra 1 bit is reserved for future purposes.

