package lib

import (
  "os"
  "encoding/csv"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

//var url = "https://s3.amazonaws.com/gocodo/usgov/hhs/icd/icd-10-cm.csv"
var url = "https://s3.amazonaws.com/gocodo/usgov/hhs/icd/icd-10-cm-2017.csv"

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.icd_10_cm",
      Version: "2017-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch(url)
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer fileReader.Close()

  csvReader := csv.NewReader(fileReader)
  if err != nil {
    return nil, err
  }

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch(url)
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