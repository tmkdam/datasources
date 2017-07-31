package lib

import (
  "os"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

var url = "https://s3.amazonaws.com/gocodo/usgov/hhs/icd/icd-9-cm.csv"

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.icd_9_cm",
      Version: "2014-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := dataloading.NewDownloader("data/", nil)
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

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
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