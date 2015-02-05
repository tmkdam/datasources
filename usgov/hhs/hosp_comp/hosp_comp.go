package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/hosp_comp/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
