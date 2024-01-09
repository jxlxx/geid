# Good Enough IDs

This is a package that generates short IDs that are "unique" enough to be used instead of UUIDs, which are ugly.

Example: `ab-ce093c1`

## Features

- :sparkles: Optional prefixes: IDs will be of the form `prefix-*`
- :sparkles: Safe to use with goroutines: New IDs require data from a counter which has a lock
- :sparkles: Safe to use with multiple machines: Each ID requires a "machine ID" int coming from the environment
- :sparkles: Optional custom epoch: Can override the default epoch of January 1st, 1970

## Install

```sh
go install github.com/jxlxx/geid@latest
```

## Configuration

Epoch is optional. The default epoch is January 1st, 1970.

```yaml
package: animals
epoch:
  year: 2024
  month: 6
  day: 14
ids:
  - name: Cat
    prefix: cat-
  - name: Dog
    prefix: dog-
  - name: something # this id will not have a prefix 
```

## Getting Started

Generate code:

```sh
geid -c config.yaml > ids.go
```

Generated code:

```go
package animals

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const CatIDPrefix = "cat-"
const DogIDPrefix = "dog-"
const somethingIDPrefix = ""

func newCatID() string {
	return fmt.Sprintf("%s%s", CatIDPrefix, generateCode())
}

func newDogID() string {
	return fmt.Sprintf("%s%s", DogIDPrefix, generateCode())
}

func newsomethingID() string {
	return fmt.Sprintf("%s%s", somethingIDPrefix, generateCode())
}

func generateCode() string {
	s := getSequenceNumber()
	t := getSecondsSinceEpoch()
	m := getMachineID()
	hex := fmt.Sprintf("%x%x%x", s, t, m)
	return hex
}

type Counter struct {
	mu sync.Mutex
	x  int
}

var counter = Counter{}

func getSequenceNumber() int {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.x = (counter.x + 1) % 4096
	return counter.x + 170 // adding 170 so that the first few IDs are prettier
}

var customEpoch = time.Date(2024, time.June, 14, 0, 0, 0, 0, time.UTC).Unix()

func getSecondsSinceEpoch() int64 {
	return time.Now().Unix() - customEpoch
}

func getMachineID() int {
	machineID := mustGetEnv("MACHINE_ID")
	i, err := strconv.Atoi(machineID)
	if err != nil {
		log.Fatalf("cannot parse machine id into int. (MACHINE_ID = %s) ", machineID)
	}
	return i
}

func mustGetEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s environment key not found.\n", key)
	}
	return v
}
```
## Examples

```sh
cat-ab-ce093c1
cat-ac-ce093c1
cat-ad-ce093c1
dog-ae-ce093c1
dog-af-ce093c1
dog-b0-ce093c1
b1-ce093c1
b2-ce093c1
b3-ce093c1
```


## Design

The implementation is loosely based on the following article: [Creating User-Facing, Short Unique IDs: What are the options? - Hwee Lin Yeo, Alexandra](https://medium.com/teamocard/creating-user-facing-short-unique-id-ids-what-are-the-options-464a19283d98)

> Twitter Snowflake, the open-source version of which is unfortunately archived, is an internal service used by Twitter for generating 64-bit unique IDs at a high scale. The IDs are made up of the components:
> - Epoch timestamp in millisecond precision — 41 bits (gives us 69 years with a custom epoch)
> - Machine id — 10 bits (thus allowing uniqueness even when we scale the short ID generator service over different nodes)
> - Sequence number — 12 bits (A local counter per machine that rolls over every 4096)
> - The extra 1 bit is reserved for future purposes.

