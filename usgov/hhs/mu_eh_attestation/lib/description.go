package lib

import (
  "os"
  "errors"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

var stageOne = "https://s3.amazonaws.com/gocodo/usgov/hhs/mu/201504+EH+Stage+1+PUF.csv"
var stageTwo = "https://s3.amazonaws.com/gocodo/usgov/hhs/mu/201504+EH+Stage+2+PUF.csv"

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.mu_eh_stage1_attestation",
      Version: "2015-01-01",
    },
    dataloading.Source{
      Name: "usgov.hhs.mu_eh_stage2_attestation",
      Version: "2015-01-01",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  var uri string
  if sourceName == "usgov.hhs.mu_eh_stage1_attestation" {
    uri = stageOne
  } else if sourceName == "usgov.hhs.mu_eh_stage2_attestation" {
    uri = stageTwo
  } else {
    return nil, errors.New("Unknown datasource name")
  }

  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
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
  var uri string
  if source.Name == "usgov.hhs.mu_eh_stage1_attestation" {
    uri = stageOne
  } else if source.Name == "usgov.hhs.mu_eh_stage2_attestation" {
    uri = stageTwo
  } else {
    return nil, errors.New("Unknown datasource name")
  }

  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
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