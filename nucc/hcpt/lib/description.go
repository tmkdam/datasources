package lib

import (
  "os"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "nucc.hcpt",
      Version: "2015-01",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  return []string{ "Code", "Type", "Classification", "Specialization", "Definition", "Notes" }, nil
}

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://www.nucc.org/images/stories/CSV/nucc_taxonomy_150.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}