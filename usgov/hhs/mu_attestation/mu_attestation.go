package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/mu_attestation/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
