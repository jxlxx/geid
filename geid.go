package geid

import (
	"fmt"
	"sync"
	"time"
)

type Prefix interface {
	Prefix() string
}

type Generator struct {
	prefix string
}

func New(p Prefix) *Generator {
	return &Generator{
		prefix: p.Prefix(),
	}
}

func (i Generator) NewID() string {
	return fmt.Sprintf("%s%s", i.prefix, generateCode())
}

var machineID = "1"

func SetMachineID(s string) {
	machineID = s
}

var epoch = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Unix()

func SetCustomEpoch(t time.Time) {
	epoch = t.Unix()
}

func getSecondsSinceEpoch() int64 {
	return time.Now().Unix() - epoch
}

func generateCode() string {
	s := getSequenceNumber()
	t := getSecondsSinceEpoch()
	hex := fmt.Sprintf("%x%x%x", s, t, machineID)
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
	return counter.x + 170
}
