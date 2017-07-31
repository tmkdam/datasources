package main

import (
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/datasources/usgov/hhs/ahrq/md_sid_2011/lib"
)

func main() {
  dataloading.CreateCmd(&lib.Description{})
}
