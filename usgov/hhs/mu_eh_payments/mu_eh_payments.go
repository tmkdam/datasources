package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/mu_eh_payments/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
