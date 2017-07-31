package lib

import (
  "io"
  "io/ioutil"
  "net/http"
  "time"
  "regexp"
  "strings"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

var downloadRegex = regexp.MustCompile(`([A-Z]+)(\d+)_OTHER_CSV`)

var shortCapMonths = map[string]string{
  "JAN": "Jan",
  "FEB": "Feb",
  "MAR": "Mar",
  "APR": "Apr",
  "MAY": "May",
  "JUN": "Jun",
  "JUL": "Jul",
  "AUG": "Aug",
  "SEP": "Sep",
  "OCT": "Oct",
  "NOV": "Nov",
  "DEC": "Dec",
}

func (d *Description) Available() ([]dataloading.Source, error) {
  resp, err := http.Get("http://www.cms.gov/Research-Statistics-Data-and-Systems/Downloadable-Public-Use-Files/Provider-of-Services/index.html")
  if err != nil {
    return nil, err
  }

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return nil, err
  }

  bodyS := string(body)

  monthlyMatches := downloadRegex.FindAllStringSubmatch(bodyS, -1)

  sources := []dataloading.Source{}

  for _, match := range monthlyMatches {
    month := match[1]
    year := match[2]

    date, err := time.Parse("Jan 2, 06", shortCapMonths[month] + " 1, " + year)
    if err != nil {
      return nil, err
    }

    version := date.Format("2006-01")

    sources = append(sources, dataloading.Source{
      Name: "usgov.hhs.pos",
      Version: version,
    })
  }

  return sources, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  sources, err := d.Available()
  if err != nil {
    return nil, err
  }

  reader, err := getFileReader(sources[0])
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

func getFileReader(source dataloading.Source) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  zipPattern := regexp.MustCompile(`\.csv$`)

  d, err := time.Parse("2006-01", source.Version)
  if err != nil {
    return nil, err
  }
  f := strings.ToUpper(d.Format("Jan06"))

  uri := "http://www.cms.gov/Research-Statistics-Data-and-Systems/Downloadable-Public-Use-Files/Provider-of-Services/Downloads/" + f + "_OTHER_CSV.zip"

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
  reader, err := getFileReader(source)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}