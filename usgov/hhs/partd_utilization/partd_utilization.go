package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/partd_utilization/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
