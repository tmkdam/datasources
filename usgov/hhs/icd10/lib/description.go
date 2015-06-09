package lib

import (
  "io"
  "regexp"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

var schema = []helpers.TabField{
                helpers.TabField{
                  Name: "code",
                  StartIndex: 7,
                  EndIndex: 14,
                },
                helpers.TabField{
                  Name: "short_description",
                  StartIndex: 17,
                  EndIndex: 77,
                },
                helpers.TabField{
                  Name: "long_description",
                  StartIndex: 78,
                  EndIndex: 1000,
                },
              }

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.icd_10_cm",
      Version: "2015-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  columns := make([]string, len(schema))

  for i, field := range schema {
    columns[i] = field.Name
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
  fileMatch := regexp.MustCompile(`icd10cm_order_2015.txt$`)
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/ICD10/Downloads/2015-code-descriptions.zip", fileMatch)
  if err != nil {
    return nil, err
  }

  return helpers.NewTabReader(reader, schema), nil
}