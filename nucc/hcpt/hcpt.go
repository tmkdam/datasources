package main

import (
	"github.com/bloomapi/dataloading"
	"github.com/bloomapi/datasources/nucc/hcpt/lib"
)

func main() {
	dataloading.CreateCmd(&lib.Description{})
}
