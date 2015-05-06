package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/mu_report/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
