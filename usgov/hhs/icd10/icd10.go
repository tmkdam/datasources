package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/icd10/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
