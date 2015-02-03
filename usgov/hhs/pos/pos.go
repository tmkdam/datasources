package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/pos/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
