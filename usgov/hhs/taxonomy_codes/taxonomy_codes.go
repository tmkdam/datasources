package main

import (
	"github.com/gocodo/bloomsource"
	"github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes/lib"
)

func main() {
	bloomsource.CreateCmd(&lib.Description{})
}
