package main

import (
	"flag"
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	// 1a. Flags definieren
	count := flag.Int("count", 1, "proverb count")

	// 1b. Flags parsen
	flag.Parse()

	// 2. Ausgabe in Schleife
	for range *count {
		var p *proverbs.Proverb = proverbs.Random()
		fmt.Println(p.Saying)
	}
}
