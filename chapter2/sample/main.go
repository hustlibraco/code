package main

import (
	"log"
	"os"

	_ "github.com/hustlibraco/code/chapter2/sample/matchers"	// 调用matchers.init()，注册rss matcher
	"github.com/hustlibraco/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("a")
}
