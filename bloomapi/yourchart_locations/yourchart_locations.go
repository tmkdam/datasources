package main

import (
	"github.com/bloomapi/dataloading"
	"github.com/bloomapi/datasources/bloomapi/yourchart_locations/lib"
)

func main() {
	dataloading.CreateCmd(&lib.Description{})
}
