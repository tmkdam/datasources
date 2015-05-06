package main

import (
	"bitbucket.org/gocodo/bloomsource"
	"bitbucket.org/gocodo/bloomsources/usgov/hhs/taxonomy_codes/lib"
)

func main() {
	bloomsource.CreateCmd(&lib.Description{})
}
