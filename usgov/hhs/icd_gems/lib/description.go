package lib

import (
  "io"
  "regexp"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

var schema9 = []helpers.TabField{
                helpers.TabField{
                  Name: "icd9",
                  StartIndex: 1,
                  EndIndex: 6,
                },
                helpers.TabField{
                  Name: "icd10",
                  StartIndex: 7,
                  EndIndex: 14,
                },
                helpers.TabField{
                  Name: "approximate",
                  StartIndex: 15,
                  EndIndex: 15,
                },
                helpers.TabField{
                  Name: "no_map",
                  StartIndex: 16,
                  EndIndex: 16,
                },
                helpers.TabField{
                  Name: "combination",
                  StartIndex: 17,
                  EndIndex: 17,
                },
                helpers.TabField{
                  Name: "scenario",
                  StartIndex: 18,
                  EndIndex: 18,
                },
                helpers.TabField{
                  Name: "choice_list",
                  StartIndex: 19,
                  EndIndex: 19,
                },
              }

var schema10 = []helpers.TabField{
                helpers.TabField{
                  Name: "icd10",
                  StartIndex: 1,
                  EndIndex: 8,
                },
                helpers.TabField{
                  Name: "icd9",
                  StartIndex: 9,
                  EndIndex: 14,
                },
                helpers.TabField{
                  Name: "approximate",
                  StartIndex: 15,
                  EndIndex: 15,
                },
                helpers.TabField{
                  Name: "no_map",
                  StartIndex: 16,
                  EndIndex: 16,
                },
                helpers.TabField{
                  Name: "combination",
                  StartIndex: 17,
                  EndIndex: 17,
                },
                helpers.TabField{
                  Name: "scenario",
                  StartIndex: 18,
                  EndIndex: 18,
                },
                helpers.TabField{
                  Name: "choice_list",
                  StartIndex: 19,
                  EndIndex: 19,
                },
              }

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.icd_10_gems",
      Version: "2017-00",
    },
    dataloading.Source{
      Name: "usgov.hhs.icd_9_gems",
      Version: "2017-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  columns := make([]string, len(schema9))

  for i, field := range schema9 {
    columns[i] = field.Name
  }

  return columns, nil
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
  url := "https://www.cms.gov/Medicare/Coding/ICD10/Downloads/2017-GEM-DC.zip"
  
  var fileMatch *regexp.Regexp
  var schema []helpers.TabField

  if source.Name == "usgov.hhs.icd_9_gems" {
    fileMatch = regexp.MustCompile(`2017_I10gem.txt$`)
    schema = schema10
  } else {
    fileMatch = regexp.MustCompile(`2017_I9gem.txt$`)
    schema = schema9
  }

  reader, err := getFileReader(url, fileMatch)
  if err != nil {
    return nil, err
  }

  return helpers.NewTabReader(reader, schema), nil
}