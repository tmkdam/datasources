package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/mu_report/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
