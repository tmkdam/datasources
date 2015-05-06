package lib

import (
  "net/http"
  "io/ioutil"
  "io"
  "encoding/csv"
  "strings"
  "regexp"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

var htmlUri = "http://nppes.viva-it.com/NPI_Files.html"

var ignore = regexp.MustCompile(`<\!\-\-(.|\n)*?\-\-\>`)
var weeklyRegex = regexp.MustCompile(`http\:\/\/.*(NPPES_Data_Dissemination_\d+_\d+_Weekly).zip`)
var monthlyRegex = regexp.MustCompile(`http\:\/\/.*(NPPES_Data_Dissemination_[a-zA-Z]+_\d+).zip`)

func (d *Description) Available() ([]bloomsource.Source, error) {
  found := []bloomsource.Source{}

  resp, err := http.Get(htmlUri)
  if err != nil {
    return nil, err
  }

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return nil, err
  }

  bodyS := string(body)

  bodyS = ignore.ReplaceAllString(bodyS, "")
  monthlyMatches := monthlyRegex.FindStringSubmatch(bodyS)
  weeklyMatches := weeklyRegex.FindAllStringSubmatch(bodyS, -1)
  
  var monthly string
  if len(monthlyMatches) > 0 {
    monthly = monthlyMatches[1]
    found = append(found, bloomsource.Source{
      "usgov.hhs.npi",
      "full-" + monthly,
      "sync",
    })
  }

  for _, weeklyMatch := range weeklyMatches {
    if len(weeklyMatch) > 0 {
      found = append(found, bloomsource.Source{
        "usgov.hhs.npi",
        "incremental-" + weeklyMatch[1],
        "upsert",
      })
    }
  }

  return found, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  sources, err := d.Available()
  if err != nil {
    return nil, err
  }
  source := sources[0]

  version := strings.Replace(source.Version, "full-", "", 1)

  uri := "http://nppes.viva-it.com/" + version + ".zip"
  zipPattern := regexp.MustCompile(`FileHeader.csv$`)

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
  version := strings.Replace(strings.Replace(source.Version, "full-", "", 1), "incremental-", "", 1)
  uri := "http://nppes.viva-it.com/" + version + ".zip"
  zipPattern := regexp.MustCompile(`\d+.csv$`)

  reader, err := getFileReader(uri, zipPattern)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}