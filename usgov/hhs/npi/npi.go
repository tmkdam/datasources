package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/npi/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
