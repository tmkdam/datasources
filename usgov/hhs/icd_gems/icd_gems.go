package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/icd_gems/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
