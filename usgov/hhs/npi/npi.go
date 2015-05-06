package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/npi/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
