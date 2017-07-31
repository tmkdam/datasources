package lib

import (
	"os"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
 	"github.com/spf13/viper"
)

type Description struct {}

var sourceToFiles = map[string]string {
	"usgov.hhs.ahrq.md_sid_2011.chgs": "MD_SID_2011_CHGS.asc",
	"usgov.hhs.ahrq.md_sid_2011.core": "MD_SID_2011_CORE.asc",
	"usgov.hhs.ahrq.md_sid_2011.dx_pr_grps": "MD_SID_2011_DX_PR_GRPS.asc",
	"usgov.hhs.ahrq.md_sid_2011.severity": "MD_SID_2011_SEVERITY.asc",
}

func (d *Description) Available() ([]dataloading.Source, error) {
	return []dataloading.Source{
			dataloading.Source{
				Name: "usgov.hhs.ahrq.md_sid_2011.chgs",
				Version: "001",
			},
			dataloading.Source{
				Name: "usgov.hhs.ahrq.md_sid_2011.core",
				Version: "001",
			},
			dataloading.Source{
				Name: "usgov.hhs.ahrq.md_sid_2011.dx_pr_grps",
				Version: "001",
			},
			dataloading.Source{
				Name: "usgov.hhs.ahrq.md_sid_2011.severity",
				Version: "001",
			},
		}, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	tabFields, err := buildSchema(sourceName)
	if err != nil {
		return nil, err
	}

	fieldNames := make([]string, len(tabFields))

	for i, field := range tabFields {
		fieldNames[i] = field.Name
	}

	return fieldNames, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
	tabFields, err := buildSchema(source.Name)
	if err != nil {
		return nil, err
	}

	cclfDir := viper.GetString("mdSid2011Dir")

	fileName := sourceToFiles[source.Name]

	path := cclfDir + "/" + fileName

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	tabReader := helpers.NewTabReader(file, tabFields)
	return tabReader, nil
}