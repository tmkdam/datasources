package lib

import (
  "io"
  "path/filepath"
  "os"
  "regexp"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
  "github.com/gocodo/bloomsources/usgov/hhs/cclf/schema"
 	"github.com/spf13/viper"
 	"errors"
)

type Description struct {}

var cclfFileName = regexp.MustCompile(`ACO\.CCLF(\d)\.(D\d\d\d\d\d\d\.T\d\d\d\d\d\d\d)`)

var sourceNames = map[string]string {
	"1": "usgov.hhs.cclf_parta",
	"2": "usgov.hhs.cclf_parta_revenue",
	"3": "usgov.hhs.cclf_parta_proc",
	"4": "usgov.hhs.cclf_parta_dx",
	"5": "usgov.hhs.cclf_partb_phy",
	"6": "usgov.hhs.cclf_partb_dme",
	"7": "usgov.hhs.cclf_partd",
	"8": "usgov.hhs.cclf_beneficiary",
	"9": "usgov.hhs.cclf_bene_xref",
}

// Inverse of sourceNames
var fileIndexes = map[string]string {
	"usgov.hhs.cclf_parta": "1",
	"usgov.hhs.cclf_parta_revenue": "2",
	"usgov.hhs.cclf_parta_proc": "3",
	"usgov.hhs.cclf_parta_dx": "4",
	"usgov.hhs.cclf_partb_phy": "5",
	"usgov.hhs.cclf_partb_dme": "6",
	"usgov.hhs.cclf_partd": "7",
	"usgov.hhs.cclf_beneficiary": "8",
	"usgov.hhs.cclf_bene_xref": "9",
}

var sourceFields = map[string]*schema.Cclf {
	"usgov.hhs.cclf_parta": schema.Parta,
	"usgov.hhs.cclf_parta_revenue": schema.PartaRevenue,
	"usgov.hhs.cclf_parta_proc": schema.PartaProc,
	"usgov.hhs.cclf_parta_dx": schema.PartaDx,
	"usgov.hhs.cclf_partb_phy": schema.PartbPhy,
	"usgov.hhs.cclf_partb_dme": schema.PartbDme,
	"usgov.hhs.cclf_partd": schema.Partd,
	"usgov.hhs.cclf_beneficiary": schema.Beneficiary,
	"usgov.hhs.cclf_bene_xref": schema.BeneXref,
}

func (d *Description) Available() ([]bloomsource.Source, error) {
	sources := []bloomsource.Source{}

	cclfDir := viper.GetString("cclfDir")

	err := filepath.Walk(cclfDir, func (path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !file.IsDir() {
				matches := cclfFileName.FindStringSubmatch(file.Name())
				if len(matches) > 0 {
					fileNum := matches[1]
					fileVersion := matches[2]
					sourceName := sourceNames[fileNum]

					if sourceName != "" {
						sources = append(sources, bloomsource.Source{
							Name: sourceName,
							Version: fileVersion,
							Action: "upsert",
						})
					}
				}
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return sources, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	fields := sourceFields[sourceName].FieldNames()
	names := make([]string, len(fields))

	for i, field := range fields {
		names[i] = field
	}

	return names, nil
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
	fileIndex := fileIndexes[source.Name]
  fileRegex := regexp.MustCompile(`ACO\.CCLF` + fileIndex + `\.` + source.Version)

	cclfDir := viper.GetString("cclfDir")
	var sourcePath string

	err := filepath.Walk(cclfDir, func (path string, file os.FileInfo, err error) error {
			if fileRegex.MatchString(file.Name()) {
				sourcePath = path
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	if sourcePath == "" {
		return nil, errors.New("Source " + source.Name + " with version " + source.Version + " not found.")
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return nil, err
	}

	return helpers.NewTabReader(file, sourceFields[source.Name].Fields(source.Version)), nil
}