package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	w *sync.WaitGroup
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	w = &sync.WaitGroup{}

	task()

	w.Wait()
}
