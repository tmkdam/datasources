package lib

import (
  "io"
  "regexp"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
)

type Description struct {}

var fields = []string{ "NPI Number1", "NPI Number2", "Pair Count", "Bene Count", "Same Day Count" }

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.phyreferral",
      Version: "2014-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  return fields, nil
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
  fileMatch := regexp.MustCompile(`Physician-Referrals-2013-2014-DAYS30.txt$`)
  reader, err := getFileReader("https://s3.amazonaws.com/gocodo/usgov/hhs/phyreferral/physician-referrals-2013-2014-days30.zip", fileMatch)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReaderNoHeaders(reader, fields)

  return zipReader, nil
}