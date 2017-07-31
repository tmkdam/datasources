package main

import (
	"github.com/bloomapi/dataloading"
	"github.com/bloomapi/datasources/usgov/hhs/taxonomy_codes/lib"
)

func main() {
	dataloading.CreateCmd(&lib.Description{})
}
