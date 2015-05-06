package lib

import (
  "os"
  "encoding/csv"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.mu_report",
      Version: "2015-00-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://www.healthit.gov/sites/default/files/mu_report.csv")
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
  path, err := downloader.Fetch("http://www.healthit.gov/sites/default/files/mu_report.csv")
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