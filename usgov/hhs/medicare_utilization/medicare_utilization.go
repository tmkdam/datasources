package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/medicare_utilization/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
