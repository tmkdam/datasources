package lib

import (
  "os"
  "errors"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "bloomapi.epic.mychart_locations",
      Version: "2015-05-27",
    },
    dataloading.Source{
      Name: "bloomapi.epic.mychart_location_states",
      Version: "2015-05-27",
    },
  }, nil
}

var urls = map[string]string {
  "bloomapi.epic.mychart_locations": "https://s3.amazonaws.com/gocodo/bloomapi/epic/epic_sites.csv",
  "bloomapi.epic.mychart_location_states": "https://s3.amazonaws.com/gocodo/bloomapi/epic/epic_site_states.csv",
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  uri, ok := urls[sourceName]
  if !ok {
    return nil, errors.New("Source not found: " + sourceName)
  }

  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

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
  uri, ok := urls[source.Name]
  if !ok {
    return nil, errors.New("Source not found: " + source.Name)
  }

  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}