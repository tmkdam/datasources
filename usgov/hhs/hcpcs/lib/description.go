package lib

import (
  "io"
  "io/ioutil"
  "regexp"
  "bytes"
  "archive/zip"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.hcpcs",
      Version: "2015-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/HCPCSReleaseCodeSets/Downloads/2015-Annual-Alpha-Numeric-HCPCS-File.zip", regexp.MustCompile(`HCPC2015_CONTR_ANWEB_v2.xlsx$`))
  if err != nil {
    return nil, err
  }

  allBytes, err := ioutil.ReadAll(reader)
  if err != nil {
  	return nil, err
  }

  bReader := bytes.NewReader(allBytes)

  zipReader, err := zip.NewReader(bReader, int64(bReader.Len()))
  if err != nil {
  	return nil, err
  }

  xlReader, err := helpers.NewXlxsReader(zipReader, "hcpc15")
  if err != nil {
  	return nil, err
  }

	return xlReader.Headers(), nil
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

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/HCPCSReleaseCodeSets/Downloads/2015-Annual-Alpha-Numeric-HCPCS-File.zip", regexp.MustCompile(`HCPC2015_CONTR_ANWEB_v2.xlsx$`))
  if err != nil {
    return nil, err
  }

  allBytes, err := ioutil.ReadAll(reader)
  if err != nil {
  	return nil, err
  }

  bReader := bytes.NewReader(allBytes)

  zipReader, err := zip.NewReader(bReader, int64(bReader.Len()))
  if err != nil {
  	return nil, err
  }

  xlReader, err := helpers.NewXlxsReader(zipReader, "hcpc15")
  if err != nil {
  	return nil, err
  }

  return xlReader, nil
}