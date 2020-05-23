package main

import (
	"fmt"
	"time"
)

const (
	date        = "2006-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:02"
	shorttime   = "15:04"
	datetime    = "2006-01-02 15:04:02"
	newdatetime = "2006/01/02 15~04~02"
	newtime     = "15~04~02"
)

func main() {
	a, _ := time.Parse("2006-01-02", "2008-01-02")
	fmt.Print(a)
}
