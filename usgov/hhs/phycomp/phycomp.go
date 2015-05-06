package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/phycomp/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
