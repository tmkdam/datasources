package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/medicare_utilization/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
