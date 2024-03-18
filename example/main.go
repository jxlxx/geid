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
