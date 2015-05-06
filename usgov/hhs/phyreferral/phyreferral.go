package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/phyreferral/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
