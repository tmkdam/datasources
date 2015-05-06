package main

import (
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsources/usgov/hhs/ahrq/md_sid_2011/lib"
)

func main() {
  bloomsource.CreateCmd(&lib.Description{})
}
