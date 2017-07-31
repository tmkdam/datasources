package lib

import (
  "net/http"
  "io/ioutil"
  "io"
  "encoding/csv"
  "errors"
  "regexp"
  "time"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

type SourceDesc struct {
  FileName string
  SourceName string
}

var rootSourceName = "usgov.hhs."

var sourceDescs = []SourceDesc{
  SourceDesc{
    FileName: "National_Downloadable_File.csv",
    SourceName: "phycomp",
  },
}

var htmlUri = "https://data.medicare.gov/data/physician-compare"
// Must Match 'CSV Flat Files &ndash; Revised</a> <div class="file-desc">Updated: Apr 16, 2015'
var lastUpdateRegexp = regexp.MustCompile(`CSV Flat Files &ndash; Revised</a> <div class="file-desc">Updated: ([\w\s\,]+)`)
// Must Match https://data.medicare.gov/views/bg9k-emty/files/ef0QT-EmZfBJUY4PSoNaTmsxMS0CJS4hq5rg67srcU4?filename=Physician_Compare_Databases.zip&content_type=application%2Fzip%3B%20charset%3Dbinary
var downloadLinkRegexp = regexp.MustCompile(`\"([^\"]+Physician_Compare_Databases[^\"]+)\"`)

func GetBody() (string, error) {
  resp, err := http.Get(htmlUri)
  if err != nil {
    return "", err
  }

  body, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    return "", err
  }

  return string(body), nil
}

func (d *Description) Available() ([]dataloading.Source, error) {
  bodyS, err := GetBody()
  if err != nil {
    return nil, err
  }

  lastUpdateMatches := lastUpdateRegexp.FindStringSubmatch(bodyS)

  t, err := time.Parse("Jan 2, 2006", lastUpdateMatches[1])
  if err != nil {
    return nil, err
  }

  sources := make([]dataloading.Source, len(sourceDescs))

  for i, sourceDesc := range sourceDescs {
    sources[i] = dataloading.Source{
      Name: rootSourceName + sourceDesc.SourceName,
      Version: t.Format(time.RFC3339),
    }
  }

  return sources, nil
}

func getFileReader(uri string, stringPattern string) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  zipPattern := regexp.MustCompile(stringPattern)

  reader, err := helpers.OpenExtractZipReader(path, zipPattern)
  if err != nil {
    return nil, err
  }

  return reader, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  bodyS, err := GetBody()
  if err != nil {
    return nil, err
  }

  downloadLinkMatches := downloadLinkRegexp.FindStringSubmatch(bodyS)
  uri := "https://data.medicare.gov" + downloadLinkMatches[1]

  var sourceDesc SourceDesc
  for _, desc := range sourceDescs {
    if rootSourceName + desc.SourceName == sourceName {
      sourceDesc = desc
      break
    }
  }

  if sourceDesc.FileName == "" {
    return nil, errors.New("Source Description Not Found")
  }

  reader, err := getFileReader(uri, sourceDesc.FileName)
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
  bodyS, err := GetBody()
  if err != nil {
    return nil, err
  }

  downloadLinkMatches := downloadLinkRegexp.FindStringSubmatch(bodyS)
  uri := "https://data.medicare.gov" + downloadLinkMatches[1]

  var sourceDesc SourceDesc
  for _, desc := range sourceDescs {
    if rootSourceName + desc.SourceName == source.Name {
      sourceDesc = desc
      break
    }
  }

  if sourceDesc.FileName == "" {
    return nil, errors.New("Source Description Not Found")
  }

  reader, err := getFileReader(uri, sourceDesc.FileName)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvReader(reader)

  return zipReader, nil
}