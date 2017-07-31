package main

import (
	"bitbucket.org/gocodo/bloomsource"
	"bitbucket.org/gocodo/bloomsources/bloomapi/yourchart_locations/lib"
)

func main() {
	bloomsource.CreateCmd(&lib.Description{})
}
