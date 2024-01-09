# Good Enough IDs

## Install

```sh
go install github.com/jxlxx/geid
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
  - name: Something # this id will not have a prefix 
```

## Generating code

```sh
geid -c config.yaml > ids.go
```

The 

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


## Example IDs

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
