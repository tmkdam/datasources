package main

import (
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsources/usgov/hhs/mu_attestation/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
