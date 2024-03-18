# Good Enough IDs

This is a package that generates short IDs that are "unique" enough to be used instead of UUIDs, which are ugly.

## Features

- :sparkles: Optional prefixes: IDs will be of the form `prefix-*`
- :sparkles: Safe to use with goroutines: New IDs require data from a counter which has a lock
- :sparkles: Safe to use with multiple machines: Can set a unique "Machine ID" for each instance
- :sparkles: Optional custom epoch: Can override the default epoch of January 1st, 1970


## Getting Started

Epoch is optional. The default epoch is January 1st, 1970. Setting the epoch to a later date means short IDs.

Setting a default prefix for IDs is optional. Or you can override the default prefix with `NewWithPrefix`.

MachineID is optional. The default machine ID is "1". This is useful if you have many instances of identical
ID generators running at the same time.

```go
package main

import (
	"fmt"
	"time"

	"github.com/jxlxx/geid"
)

func main() {
	t := time.Now()
	geid.SetCustomEpoch(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC))

	geid.SetDefaultPrefix("cat-")

	geid.SetMachineID("1")

	for i := 0; i < 10; i++ {
		fmt.Println(geid.New())
	}
	fmt.Println(geid.NewWithPrefix("dog-"))
}
```

Output:

```sh
cat-ab384031
cat-ac384031
cat-ad384031
cat-ae384031
cat-af384031
cat-b0384031
cat-b1384031
cat-b2384031
cat-b3384031
cat-b4384031
dog-b5384031
```

## Design

The implementation is loosely based on the following article: [Creating User-Facing, Short Unique IDs: What are the options? - Hwee Lin Yeo, Alexandra](https://medium.com/teamocard/creating-user-facing-short-unique-id-ids-what-are-the-options-464a19283d98)

> Twitter Snowflake, the open-source version of which is unfortunately archived, is an internal service used by Twitter for generating 64-bit unique IDs at a high scale. The IDs are made up of the components:
> - Epoch timestamp in millisecond precision — 41 bits (gives us 69 years with a custom epoch)
> - Machine id — 10 bits (thus allowing uniqueness even when we scale the short ID generator service over different nodes)
> - Sequence number — 12 bits (A local counter per machine that rolls over every 4096)
> - The extra 1 bit is reserved for future purposes.

