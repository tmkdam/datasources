package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/npi/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
