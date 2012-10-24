package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
	"time"
)

func main() {
	now := time.Now().Format("3:04")
	avg := sigar.LoadAverage{}
	avg.Get()
	mem := sigar.Mem{}
	mem.Get()

	fmt.Printf("%s | %.2f, %.2f, %.2f | %s\n",
		now,
		avg.One, avg.Five, avg.Fifteen,
		sigar.FormatSize(mem.ActualFree))
}
