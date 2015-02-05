package lib

import (
  "net/http"
  "io/ioutil"
  "io"
  "encoding/csv"
  "errors"
  "regexp"
  "time"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
)

type Description struct {}

type SourceDesc struct {
  FileName string
  SourceName string
}

var rootSourceName = "usgov.hhs.hosp_comp."

var sourceDescs = []SourceDesc{
  SourceDesc{
    FileName: "HOSPITAL_QUARTERLY_QUALITYMEASURE_IPFQR_HOSPITAL.csv",
    SourceName: "hosp_quart_qm_ipfqr_hospital",
  },
  SourceDesc{
    FileName: "HOSPITAL_QUARTERLY_QUALITYMEASURE_IPFQR_NATIONAL.csv",
    SourceName: "hosp_quart_qm_ipfqr_national",
  },
  SourceDesc{
    FileName: "HOSPITAL_QUARTERLY_QUALITYMEASURE_IPFQR_STATE.csv",
    SourceName: "hosp_quart_qm_ipfqr_state",
  },
  SourceDesc{
    FileName: "Footnote Crosswalk.csv",
    SourceName: "hqi_footnotes",
  },
  SourceDesc{
    FileName: "Hospital General Information.csv",
    SourceName: "hqi_hosp",
  },
  SourceDesc{
    FileName: "Healthcare Associated Infections - Hospital.csv",
    SourceName: "hqi_hosp_hai",
  },
  SourceDesc{
    FileName: "HCAHPS - Hospital.csv",
    SourceName: "hqi_hosp_hcahps",
  },
  SourceDesc{
    FileName: "Outpatient Imaging Efficiency - Hospital.csv",
    SourceName: "hqi_hosp_img",
  },
  SourceDesc{
    FileName: "Medicare Volume - Hospital.csv",
    SourceName: "hqi_hosp_mv",
  },
  SourceDesc{
    FileName: "Readmissions Complications and Deaths - Hospital.csv",
    SourceName: "hqi_hosp_readm_death",
  },
  SourceDesc{
    FileName: "Medicare Hospital Spending per Patient - Hospital.csv",
    SourceName: "hqi_hosp_spp",
  },
  SourceDesc{
    FileName: "Structural Measures - Hospital.csv",
    SourceName: "hqi_hosp_struct",
  },
  SourceDesc{
    FileName: "Timely and Effective Care - Hospital.csv",
    SourceName: "hqi_hosp_timelyeffective",
  },
  SourceDesc{
    FileName: "Healthcare Associated Infections - National.csv",
    SourceName: "hqi_national_hai",
  },
  SourceDesc{
    FileName: "HCAHPS - National.csv",
    SourceName: "hqi_national_hcahps",
  },
  SourceDesc{
    FileName: "Outpatient Imaging Efficiency - National.csv",
    SourceName: "hqi_national_img_avg",
  },
  SourceDesc{
    FileName: "Medicare Volume - National.csv",
    SourceName: "hqi_national_mv",
  },
  SourceDesc{
    FileName: "Readmissions Complications and Deaths - National.csv",
    SourceName: "hqi_national_readm_death",
  },
  SourceDesc{
    FileName: "Medicare Hospital Spending per Patient - National.csv",
    SourceName: "hqi_national_spp",
  },
  SourceDesc{
    FileName: "Timely and Effective Care - National.csv",
    SourceName: "hqi_national_timelyeffective",
  },
  SourceDesc{
    FileName: "Outpatient Procedures - Volume.csv",
    SourceName: "hqi_op_proc_vol",
  },
  SourceDesc{
    FileName: "Healthcare Associated Infections - State.csv",
    SourceName: "hqi_state_hai",
  },
  SourceDesc{
    FileName: "HCAHPS - State.csv",
    SourceName: "hqi_state_hcahps",
  },
  SourceDesc{
    FileName: "Outpatient Imaging Efficiency - State.csv",
    SourceName: "hqi_state_img_avg",
  },
  SourceDesc{
    FileName: "Medicare Volume - State.csv",
    SourceName: "hqi_state_mv",
  },
  SourceDesc{
    FileName: "Readmissions Complications and Deaths - State.csv",
    SourceName: "hqi_state_readm_death",
  },
  SourceDesc{
    FileName: "Medicare Hospital Spending per Patient - State.csv",
    SourceName: "hqi_state_spp",
  },
  SourceDesc{
    FileName: "Timely and Effective Care - State.csv",
    SourceName: "hqi_state_timelyeffective",
  },
  SourceDesc{
    FileName: "hvbp_ami_.+.csv",
    SourceName: "hvbp_ami",
  },
  SourceDesc{
    FileName: "hvbp_hai_.+.csv",
    SourceName: "hvbp_hai",
  },
  SourceDesc{
    FileName: "hvbp_hcahps_.+.csv",
    SourceName: "hvbp_hcahps",
  },
  SourceDesc{
    FileName: "hvbp_hf_.+.csv",
    SourceName: "hvbp_hf",
  },
  SourceDesc{
    FileName: "hvbp_outcome_.+.csv",
    SourceName: "hvbp_outcome",
  },
  SourceDesc{
    FileName: "hvbp_pn_.+.csv",
    SourceName: "hvbp_pn",
  },
  SourceDesc{
    FileName: "hvbp_quarters.csv",
    SourceName: "hvbp_quarters",
  },
  SourceDesc{
    FileName: "hvbp_scip_.+.csv",
    SourceName: "hvbp_scip",
  },
  SourceDesc{
    FileName: "HVBP_TPS_.+.csv",
    SourceName: "hvbp_tps",
  },
  SourceDesc{
    FileName: "Measure Dates.csv",
    SourceName: "measure_dates",
  },
  SourceDesc{
    FileName: "Medicare Hospital Spending by Claim.csv",
    SourceName: "medicare_spending_by_claim",
  },
  SourceDesc{
    FileName: "READMISSION REDUCTION.csv",
    SourceName: "hqi_readm_reduction",
  },
}

var htmlUri = "https://data.medicare.gov/data/hospital-compare"
var lastUpdateRegexp = regexp.MustCompile(`Updated\:\s([\w\s\,]+)`)
// Must Match /views/bg9k-emty/files/XFcQmXqywcdnwlne2L_VErOl8Hva8YTgy3d3SIVdaR4?filename=Hospital_Revised_Flatfiles.zip&content_type=application%2Fzip%3B%20charset%3Dbinary
var downloadLinkRegexp = regexp.MustCompile(`\"([^\"]+Hospital_Revised_Flatfiles[^\"]+)\"`)

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

func (d *Description) Available() ([]bloomsource.Source, error) {
  bodyS, err := GetBody()
  if err != nil {
    return nil, err
  }

  lastUpdateMatches := lastUpdateRegexp.FindStringSubmatch(bodyS)

  t, err := time.Parse("Jan 2, 2006", lastUpdateMatches[1])
  if err != nil {
    return nil, err
  }

  sources := make([]bloomsource.Source, len(sourceDescs))

  for i, sourceDesc := range sourceDescs {
    sources[i] = bloomsource.Source{
      Name: rootSourceName + sourceDesc.SourceName,
      Version: t.Format(time.RFC3339),
    }
  }

  return sources, nil
}

func getFileReader(uri string, stringPattern string) (io.Reader, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
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

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
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