package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg conc.WaitGroup
	defer wg.Wait()
	startTheThing(&wg)
}

func startTheThing(wg *conc.WaitGroup) {
	wg.Go(func() {
		n := rand.Intn(30)
		fmt.Println("will sleep ", n, " second")
		time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	})
}
