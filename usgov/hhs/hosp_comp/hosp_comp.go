package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/hosp_comp/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
