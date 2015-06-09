package lib

import (
  "io"
  "io/ioutil"
  "regexp"
  "bytes"
  "archive/zip"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.icd_9_cm",
      Version: "2014-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/ICD9ProviderDiagnosticCodes/Downloads/ICD-9-CM-v32-master-descriptions.zip", regexp.MustCompile(`CMS32_DESC_LONG_SHORT_DX.xlsx$`))
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

  xlReader, err := helpers.NewXlxsReader(zipReader, "Sheet1")
  if err != nil {
  	return nil, err
  }

	return xlReader.Headers(), nil
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
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/ICD9ProviderDiagnosticCodes/Downloads/ICD-9-CM-v32-master-descriptions.zip", regexp.MustCompile(`CMS32_DESC_LONG_SHORT_DX.xlsx$`))
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

  xlReader, err := helpers.NewXlxsReader(zipReader, "Sheet1")
  if err != nil {
  	return nil, err
  }

  return xlReader, nil
}