package lib

import (
  "io"
  "regexp"
  "encoding/csv"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.medicare_utilization",
      Version: "2012-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  fileMatch := regexp.MustCompile(`Medicare-Physician-and-Other-Supplier-PUF-CY2012.txt$`)
  reader, err := getFileReader("http://download.cms.gov/Research-Statistics-Data-and-Systems/Statistics-Trends-and-Reports/Medicare-Provider-Charge-Data/Downloads/Medicare-Physician-and-Other-Supplier-PUF-CY2012.zip?agree=yes&next=Accept", fileMatch)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(reader)
  if err != nil {
    return nil, err
  }

  csvReader.Comma = '\t'

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func getFileReader(uri string, zipPattern *regexp.Regexp) (io.Reader, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := helpers.OpenExtractZipReader(path, zipPattern)
  if err != nil {
    return nil, err
  }

  return reader, nil
}

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
  fileMatch := regexp.MustCompile(`Medicare-Physician-and-Other-Supplier-PUF-CY2012.txt$`)
  reader, err := getFileReader("http://download.cms.gov/Research-Statistics-Data-and-Systems/Statistics-Trends-and-Reports/Medicare-Provider-Charge-Data/Downloads/Medicare-Physician-and-Other-Supplier-PUF-CY2012.zip?agree=yes&next=Accept", fileMatch)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvTabReader(reader)

  return zipReader, nil
}