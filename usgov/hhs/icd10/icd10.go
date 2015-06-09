package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/icd10/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
