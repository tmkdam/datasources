package lib

import (
  "os"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.mu_attestation",
      Version: "2015-01-20",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://s3.amazonaws.com/gocodo/usgov/hhs/20150120+EH+Attestation+Summary+-+Final+-+Stage+1.csv")
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
  path, err := downloader.Fetch("http://s3.amazonaws.com/gocodo/usgov/hhs/20150120+EH+Attestation+Summary+-+Final+-+Stage+1.csv")
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