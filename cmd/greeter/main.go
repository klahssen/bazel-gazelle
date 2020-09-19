package main

import (
	"os"

	"github.com/klahssen/bazel-gazelle/pkg/greet"
)

func main() {
	greet.Greet(os.Stdout, "John Doe")
}
