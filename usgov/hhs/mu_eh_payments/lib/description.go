package lib

import (
  "io"
  "regexp"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.mu_eh_payments",
      Version: "2015-01-01",
    },
  }, nil
}

func getFileReader(uri string, zipPattern *regexp.Regexp) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
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

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  uri := "http://www.cms.gov/Regulations-and-Guidance/Legislation/EHRIncentivePrograms/Downloads/EH_ProvidersPaidByEHRProgram_Dec2014.zip"

  zipPattern := regexp.MustCompile(`EH_ProvidersPaidByEHR_12_2014_FINAL.csv$`)

  reader, err := getFileReader(uri, zipPattern)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(reader)
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
  uri := "http://www.cms.gov/Regulations-and-Guidance/Legislation/EHRIncentivePrograms/Downloads/EH_ProvidersPaidByEHRProgram_Dec2014.zip"
  zipPattern := regexp.MustCompile(`EH_ProvidersPaidByEHR_12_2014_FINAL.csv$`)

  reader, err := getFileReader(uri, zipPattern)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}