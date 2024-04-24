package main

import (
	"flag"
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1. Flags definieren
	count := flag.Int("count", 1, "proverb count")

	// 2. Flags parsen
	flag.Parse()

	// 3. Ausgabe in Schleife
	for range *count {
		p := proverbs.Random()
		fmt.Println(p.Saying)
	}
}
