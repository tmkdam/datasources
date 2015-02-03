package lib

import (
  "io"
  "regexp"
  "errors"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.medicare_specialty_codes",
      Version: "2014-09",
    },
    bloomsource.Source{
      Name: "usgov.hhs.nucc_taxonomy_codes",
      Version: "2014-09",
    },
    bloomsource.Source{
      Name: "usgov.hhs.medicare_specialty_to_nucc_taxonomy",
      Version: "2014-09",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  switch(sourceName){
  case "usgov.hhs.medicare_specialty_codes":
    return []string{ "medicare_specialty_code", "medicare_specialty_description" }, nil
  case "usgov.hhs.nucc_taxonomy_codes":
    return []string{ "taxonomy_code", "taxonomy_description" }, nil
  case "usgov.hhs.medicare_specialty_to_nucc_taxonomy":
    return []string{ "medicare_specialty_code", "nucc_taxonomy_code" }, nil
  }

  return nil, errors.New("Invalid Source")
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
  var zipPattern *regexp.Regexp
  switch(source.Name){
  case "usgov.hhs.medicare_specialty_codes":
    zipPattern = regexp.MustCompile(`medicare_specialty_codes.csv$`)
  case "usgov.hhs.nucc_taxonomy_codes":
    zipPattern = regexp.MustCompile(`nucc_taxonomy_codes.csv$`)
  case "usgov.hhs.medicare_specialty_to_nucc_taxonomy":
    zipPattern = regexp.MustCompile(`medicare_specialty_to_nucc_taxonomy.csv$`)
  }

  reader, err := getFileReader("http://s3.amazonaws.com/gocodo/medicare-specialty-codes-2014-09.zip", zipPattern)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}