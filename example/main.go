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
