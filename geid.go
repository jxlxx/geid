package geid

import (
	"fmt"
	"sync"
	"time"
)

func New() string {
	return fmt.Sprintf("%s%s", defaultPrefix, generateCode())
}

func NewWithPrefix(p string) string {
	return fmt.Sprintf("%s%s", p, generateCode())
}

var machineID = ""

func SetMachineID(s string) {
	machineID = s
}

var defaultPrefix = ""

func SetDefaultPrefix(s string) {
	defaultPrefix = s
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
