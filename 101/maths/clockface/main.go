package main

import (
	"os"
	"time"

	clockface "hello/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
