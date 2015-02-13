package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/ahrq/md_sid_2011/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
