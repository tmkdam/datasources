package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/hosp_comp/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
