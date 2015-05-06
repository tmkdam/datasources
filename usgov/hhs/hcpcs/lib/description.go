// http://www.cms.gov/Medicare/Coding/HCPCSReleaseCodeSets/Downloads/2015-Annual-Alpha-Numeric-HCPCS-File.zip
// HCPC2015_CONTR_ANWEB_v2.txt

package lib

import (
  "io"
  "regexp"
  "bitbucket.org/gocodo/bloomsource"
  "bitbucket.org/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.hcpcs",
      Version: "2015-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	fieldNames := make([]string, len(fields))
	for i, tField := range fields {
		fieldNames[i] = tField.Name
	}
	return fieldNames, nil
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
  reader, err := getFileReader("http://www.cms.gov/Medicare/Coding/HCPCSReleaseCodeSets/Downloads/2015-Annual-Alpha-Numeric-HCPCS-File.zip", regexp.MustCompile(`HCPC2015_CONTR_ANWEB_v2.txt$`))
  if err != nil {
    return nil, err
  }

  tabReader := helpers.NewTabReader(reader, fields)

  return tabReader, nil
}

var fields = []helpers.TabField{
	helpers.TabField{
		Name: "Healthcare Common Procedure Coding System Code",
		StartIndex: 1,
		EndIndex: 5,
	},
	helpers.TabField{
		Name: "HCPCS Code Redefinition Group",
		StartIndex: 1,
		EndIndex: 5,
	},
	helpers.TabField{
		Name: "HCPCS Modifier Code",
		StartIndex: 4,
		EndIndex: 5,
	},
	helpers.TabField{
		Name: "HCPCS Sequence Number",
		StartIndex: 6,
		EndIndex: 10,
	},
	helpers.TabField{
		Name: "HCPCS Record Identification Code",
		StartIndex: 11,
		EndIndex: 11,
	},
	helpers.TabField{
		Name: "HCPCS Long Description",
		StartIndex: 12,
		EndIndex: 91,
	},
	helpers.TabField{
		Name: "HCPCS Short Description",
		StartIndex: 92,
		EndIndex: 119,
	},
	helpers.TabField{
		Name: "HCPCS Pricing Indicator Code",
		StartIndex: 120,
		EndIndex: 121,
	},
	helpers.TabField{
		Name: "HCPCS Multiple Pricing Indicator Code",
		StartIndex: 128,
		EndIndex: 128,
	},
	helpers.TabField{
		Name: "HCPCS Coverage Issues Manual Reference Section Number",
		StartIndex: 129,
		EndIndex: 134,
	},
	helpers.TabField{
		Name: "HCPCS Medicare Carriers Manual Reference Section Number",
		StartIndex: 147,
		EndIndex: 154,
	},
	helpers.TabField{
		Name: "HCPCS Statute Number",
		StartIndex: 171,
		EndIndex: 180,
	},
	helpers.TabField{
		Name: "HCPCS Lab Certification Code",
		StartIndex: 181,
		EndIndex: 183,
	},
	helpers.TabField{
		Name: "HCPCS Cross Reference Code",
		StartIndex: 205,
		EndIndex: 209,
	},
	helpers.TabField{
		Name: "HCPCS Coverage Code",
		StartIndex: 230,
		EndIndex: 230,
	},
	helpers.TabField{
		Name: "HCPCS ASC Payment Group Code",
		StartIndex: 131,
		EndIndex: 132,
	},
	helpers.TabField{
		Name: "HCPCS ASC Payment Group Effective Date",
		StartIndex: 233,
		EndIndex: 240,
	},
	helpers.TabField{
		Name: "HCPCS MOG Payment Group Code",
		StartIndex: 241,
		EndIndex: 243,
	},
	helpers.TabField{
		Name: "HCPCS MOG Payment Policy Indicator",
		StartIndex: 244,
		EndIndex: 244,
	},
	helpers.TabField{
		Name: "HCPCS MOG Effective Date",
		StartIndex: 245,
		EndIndex: 252,
	},
	helpers.TabField{
		Name: "HCPCS Processing Note Number",
		StartIndex: 253,
		EndIndex: 256,
	},
	helpers.TabField{
		Name: "HCPCS Berenson-Eggers Type Of Service Code",
		StartIndex: 257,
		EndIndex: 259,
	},
	helpers.TabField{
		Name: "HCPCS Type Of Service Code",
		StartIndex: 261,
		EndIndex: 261,
	},
	helpers.TabField{
		Name: "HCPCS Anesthesia Base Unit Quantity",
		StartIndex: 266,
		EndIndex: 268,
	},
	helpers.TabField{
		Name: "HCPCS Code Added Date",
		StartIndex: 269,
		EndIndex: 276,
	},
	helpers.TabField{
		Name: "HCPCS Action Effective Date",
		StartIndex: 277,
		EndIndex: 284,
	},
	helpers.TabField{
		Name: "HCPCS Termination Date",
		StartIndex: 285,
		EndIndex: 292,
	},
	helpers.TabField{
		Name: "HCPCS Action Code",
		StartIndex: 293,
		EndIndex: 293,
	},
}