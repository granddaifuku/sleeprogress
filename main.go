package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	sec float64
	h   = flag.Float64("h", 0, "Hour")
	m   = flag.Float64("m", 0, "Minute")
	ms  = flag.Float64("ms", 0, "Millisecond")
)

func main() {
	now := time.Now()
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Invalid length of argument: required 1")
	}

	sec, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	sum := int(*h*3600*1000 + *m*60*1000 + sec*1000 + +*ms)

	bar := pb.StartNew(sum)

	for i := 0; i < sum; i++ {
		bar.Increment()
		if i%100 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}
	rem := time.Duration(sum % 100)
	time.Sleep(rem * time.Millisecond)

	bar.Finish()
	fmt.Printf("Elapsed: %vms\n", time.Since(now).Milliseconds())
}
