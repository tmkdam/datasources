package lib

import (
  "io"
  "io/ioutil"
  "os"
  "regexp"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
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
	"8": "usgov.hhs.cclf_demographics",
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
	"usgov.hhs.cclf_demographics": "8",
	"usgov.hhs.cclf_bene_xref": "9",
}

var sourceFields = map[string][]helpers.TabField{
	"usgov.hhs.cclf_parta": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Provider OSCAR Number",
			StartIndex: 14,
			EndIndex: 19,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 20,
			EndIndex: 30,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 31,
			EndIndex: 32,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 33,
			EndIndex: 42,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 43,
			EndIndex: 52,
		},
		helpers.TabField{
			Name: "Claim Bill Facility Type Code",
			StartIndex: 53,
			EndIndex: 53,
		},
		helpers.TabField{
			Name: "Claim Bill Classification Code",
			StartIndex: 54,
			EndIndex: 54,
		},
		helpers.TabField{
			Name: "Principal Diagnosis Code",
			StartIndex: 55,
			EndIndex: 61,
		},
		helpers.TabField{
			Name: "Admitting Diagnosis Code",
			StartIndex: 62,
			EndIndex: 68,
		},
		helpers.TabField{
			Name: "Claim Medicare Non-Payment Reason Code",
			StartIndex: 69,
			EndIndex: 70,
		},
		helpers.TabField{
			Name: "Claim Payment Amount",
			StartIndex: 71,
			EndIndex: 87,
		},
		helpers.TabField{
			Name: "Claim NCH Primary Payer Code",
			StartIndex: 88,
			EndIndex: 88,
		},
		helpers.TabField{
			Name: "Federal Information Processing Standards (FIPS) State Code",
			StartIndex: 89,
			EndIndex: 90,
		},
		helpers.TabField{
			Name: "Beneficiary Patient Status Code",
			StartIndex: 91,
			EndIndex: 92,
		},
		helpers.TabField{
			Name: "Diagnosis Related Group Code",
			StartIndex: 93,
			EndIndex: 96,
		},
		helpers.TabField{
			Name: "Claim Outpatient Service Type Code",
			StartIndex: 97,
			EndIndex: 97,
		},
		helpers.TabField{
			Name: "Facility Provider NPI Number",
			StartIndex: 98,
			EndIndex: 107,
		},
		helpers.TabField{
			Name: "Operating Provider NPI Number",
			StartIndex: 108,
			EndIndex: 117,
		},
		helpers.TabField{
			Name: "Attending Provider NPI Number",
			StartIndex: 118,
			EndIndex: 127,
		},
		helpers.TabField{
			Name: "Other Provider NPI Number",
			StartIndex: 128,
			EndIndex: 137,
		},
		helpers.TabField{
			Name: "Claim Adjustment Type Code",
			StartIndex: 138,
			EndIndex: 139,
		},
		helpers.TabField{
			Name: "Claim Effective Date",
			StartIndex: 140,
			EndIndex: 149,
		},
		helpers.TabField{
			Name: "Claim IDR Load Date",
			StartIndex: 150,
			EndIndex: 159,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 160,
			EndIndex: 170,
		},
		helpers.TabField{
			Name: "Claim Admission Type Code",
			StartIndex: 171,
			EndIndex: 172,
		},
		helpers.TabField{
			Name: "Claim Admission Source Code",
			StartIndex: 173,
			EndIndex: 174,
		},
		helpers.TabField{
			Name: "Claim Bill Frequency Code",
			StartIndex: 175,
			EndIndex: 175,
		},
		helpers.TabField{
			Name: "Claim Query Code",
			StartIndex: 176,
			EndIndex: 176,
		},
	},
	"usgov.hhs.cclf_parta_revenue": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Claim Line Number",
			StartIndex: 14,
			EndIndex: 23,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 24,
			EndIndex: 34,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 35,
			EndIndex: 36,
		},
		helpers.TabField{
			Name: "Claim Line From Date",
			StartIndex: 37,
			EndIndex: 46,
		},
		helpers.TabField{
			Name: "Claim Line Thru Date",
			StartIndex: 47,
			EndIndex: 56,
		},
		helpers.TabField{
			Name: "Product Revenue Center Code",
			StartIndex: 57,
			EndIndex: 60,
		},
		helpers.TabField{
			Name: "Claim Line Institutional Revenue Center Date",
			StartIndex: 61,
			EndIndex: 70,
		},
		helpers.TabField{
			Name: "HCPCS Code",
			StartIndex: 71,
			EndIndex: 75,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 76,
			EndIndex: 86,
		},
		helpers.TabField{
			Name: "Provider OSCAR Number",
			StartIndex: 87,
			EndIndex: 92,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 93,
			EndIndex: 102,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 103,
			EndIndex: 112,
		},
		helpers.TabField{
			Name: "Claim Line Allowed Unit Quantity",
			StartIndex: 113,
			EndIndex: 123,
		},
		helpers.TabField{
			Name: "Claim Line Covered Paid Amount",
			StartIndex: 124,
			EndIndex: 140,
		},
		helpers.TabField{
			Name: "HCPCS First Modifier Code",
			StartIndex: 141,
			EndIndex: 142,
		},
		helpers.TabField{
			Name: "HCPCS Second Modifier Code",
			StartIndex: 143,
			EndIndex: 144,
		},
		helpers.TabField{
			Name: "HCPCS Third Modifier Code",
			StartIndex: 145,
			EndIndex: 146,
		},
		helpers.TabField{
			Name: "HCPCS Fourth Modifier Code",
			StartIndex: 147,
			EndIndex: 148,
		},
		helpers.TabField{
			Name: "HCPCS Fifth Modifier Code",
			StartIndex: 149,
			EndIndex: 150,
		},
	},
	"usgov.hhs.cclf_parta_proc": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 14,
			EndIndex: 24,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 25,
			EndIndex: 26,
		},
		helpers.TabField{
			Name: "Claim Value Sequence Number",
			StartIndex: 27,
			EndIndex: 28,
		},
		helpers.TabField{
			Name: "Procedure Code",
			StartIndex: 29,
			EndIndex: 35,
		},
		helpers.TabField{
			Name: "Procedure Performed Date",
			StartIndex: 36,
			EndIndex: 45,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 46,
			EndIndex: 56,
		},
		helpers.TabField{
			Name: "Provider OSCAR Number",
			StartIndex: 57,
			EndIndex: 62,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 63,
			EndIndex: 72,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 73,
			EndIndex: 82,
		},
	},
	"usgov.hhs.cclf_parta_dx": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 14,
			EndIndex: 24,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 25,
			EndIndex: 26,
		},
		helpers.TabField{
			Name: "Claim Product Type Code",
			StartIndex: 27,
			EndIndex: 27,
		},
		helpers.TabField{
			Name: "Claim Value Sequence Number",
			StartIndex: 28,
			EndIndex: 29,
		},
		helpers.TabField{
			Name: "Diagnosis Code",
			StartIndex: 30,
			EndIndex: 36,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 37,
			EndIndex: 47,
		},
		helpers.TabField{
			Name: "Provider OSCAR Number",
			StartIndex: 48,
			EndIndex: 53,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 54,
			EndIndex: 63,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 64,
			EndIndex: 73,
		},
		helpers.TabField{
			Name: "Claim Present-on-Admission Indicator",
			StartIndex: 74,
			EndIndex: 80,
		},
	},
	"usgov.hhs.cclf_partb_phy": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Claim Line Number",
			StartIndex: 14,
			EndIndex: 23,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 24,
			EndIndex: 34,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 35,
			EndIndex: 36,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 37,
			EndIndex: 46,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 47,
			EndIndex: 56,
		},
		helpers.TabField{
			Name: "Rendering Provider Type Code",
			StartIndex: 57,
			EndIndex: 59,
		},
		helpers.TabField{
			Name: "Rendering Provider FIPS State Code",
			StartIndex: 60,
			EndIndex: 61,
		},
		helpers.TabField{
			Name: "Claim-Line Provider Specialty Code",
			StartIndex: 62,
			EndIndex: 63,
		},
		helpers.TabField{
			Name: "Claim Federal Type Service Code",
			StartIndex: 64,
			EndIndex: 64,
		},
		helpers.TabField{
			Name: "Claim Place of Service Code",
			StartIndex: 65,
			EndIndex: 66,
		},
		helpers.TabField{
			Name: "Claim Line From Date",
			StartIndex: 67,
			EndIndex: 76,
		},
		helpers.TabField{
			Name: "Claim Line Thru Date",
			StartIndex: 77,
			EndIndex: 86,
		},
		helpers.TabField{
			Name: "HCPCS Code",
			StartIndex: 87,
			EndIndex: 91,
		},
		helpers.TabField{
			Name: "Claim Line Covered Paid Amount",
			StartIndex: 92,
			EndIndex: 106,
		},
		helpers.TabField{
			Name: "Claim Primary Payer Code",
			StartIndex: 107,
			EndIndex: 107,
		},
		helpers.TabField{
			Name: "Diagnosis Code",
			StartIndex: 108,
			EndIndex: 114,
		},
		helpers.TabField{
			Name: "Claim Provider Tax Number",
			StartIndex: 115,
			EndIndex: 124,
		},
		helpers.TabField{
			Name: "Rendering Provider NPI Number",
			StartIndex: 125,
			EndIndex: 134,
		},
		helpers.TabField{
			Name: "Claim Carrier Payment Denial Code",
			StartIndex: 135,
			EndIndex: 136,
		},
		helpers.TabField{
			Name: "Claim-Line Processing Indicator Code",
			StartIndex: 137,
			EndIndex: 138,
		},
		helpers.TabField{
			Name: "Claim Adjustment Type Code",
			StartIndex: 139,
			EndIndex: 140,
		},
		helpers.TabField{
			Name: "Claim Effective Date",
			StartIndex: 141,
			EndIndex: 150,
		},
		helpers.TabField{
			Name: "Claim IDR Load Date",
			StartIndex: 151,
			EndIndex: 160,
		},
		helpers.TabField{
			Name: "Claim Control Number",
			StartIndex: 161,
			EndIndex: 200,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 201,
			EndIndex: 211,
		},
		helpers.TabField{
			Name: "Claim Line Allowed Charges Amount",
			StartIndex: 212,
			EndIndex: 228,
		},
		helpers.TabField{
			Name: "Claim Line Allowed Unit Quantity",
			StartIndex: 229,
			EndIndex: 239,
		},
		helpers.TabField{
			Name: "HCPCS First Modifier Code",
			StartIndex: 240,
			EndIndex: 241,
		},
		helpers.TabField{
			Name: "HCPCS Second Modifier Code",
			StartIndex: 242,
			EndIndex: 243,
		},
		helpers.TabField{
			Name: "HCPCS Third Modifier Code",
			StartIndex: 244,
			EndIndex: 245,
		},
		helpers.TabField{
			Name: "HCPCS Fourth Modifier Code",
			StartIndex: 246,
			EndIndex: 247,
		},
		helpers.TabField{
			Name: "HCPCS Fifth Modifier Code",
			StartIndex: 248,
			EndIndex: 249,
		},
		helpers.TabField{
			Name: "Claim Disposition Code",
			StartIndex: 250,
			EndIndex: 251,
		},
		helpers.TabField{
			Name: "Claim Diagnosis First Code",
			StartIndex: 252,
			EndIndex: 258,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Second Code",
			StartIndex: 259,
			EndIndex: 265,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Third Code",
			StartIndex: 266,
			EndIndex: 272,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Fourth Code",
			StartIndex: 273,
			EndIndex: 279,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Fifth Code",
			StartIndex: 280,
			EndIndex: 286,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Sixth Code",
			StartIndex: 287,
			EndIndex: 293,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Seventh Code",
			StartIndex: 294,
			EndIndex: 300,
		},
		helpers.TabField{
			Name: "Claim Diagnosis Eighth Code",
			StartIndex: 301,
			EndIndex: 307,
		},
	},
	"usgov.hhs.cclf_partb_dme": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Claim Line Number",
			StartIndex: 14,
			EndIndex: 23,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 24,
			EndIndex: 34,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 35,
			EndIndex: 36,
		},
		helpers.TabField{
			Name: "Claim From Date",
			StartIndex: 37,
			EndIndex: 46,
		},
		helpers.TabField{
			Name: "Claim Thru Date",
			StartIndex: 47,
			EndIndex: 56,
		},
		helpers.TabField{
			Name: "Claim Federal Type Service Code",
			StartIndex: 57,
			EndIndex: 57,
		},
		helpers.TabField{
			Name: "Claim Place of Service Code",
			StartIndex: 58,
			EndIndex: 59,
		},
		helpers.TabField{
			Name: "Claim Line From Date",
			StartIndex: 60,
			EndIndex: 69,
		},
		helpers.TabField{
			Name: "Claim Line Thru Date",
			StartIndex: 70,
			EndIndex: 79,
		},
		helpers.TabField{
			Name: "HCPCS Code",
			StartIndex: 80,
			EndIndex: 84,
		},
		helpers.TabField{
			Name: "Claim Line Covered Paid Amount",
			StartIndex: 85,
			EndIndex: 99,
		},
		helpers.TabField{
			Name: "Claim Primary Payer Code",
			StartIndex: 100,
			EndIndex: 100,
		},
		helpers.TabField{
			Name: "Pay-to Provider NPI Number",
			StartIndex: 101,
			EndIndex: 110,
		},
		helpers.TabField{
			Name: "Ordering Provider NPI Number",
			StartIndex: 111,
			EndIndex: 120,
		},
		helpers.TabField{
			Name: "Claim Carrier Payment Denial Code",
			StartIndex: 121,
			EndIndex: 122,
		},
		helpers.TabField{
			Name: "Claim Processing Indicator Code",
			StartIndex: 123,
			EndIndex: 124,
		},
		helpers.TabField{
			Name: "Claim Adjustment Type Code",
			StartIndex: 125,
			EndIndex: 126,
		},
		helpers.TabField{
			Name: "Claim Effective Date",
			StartIndex: 127,
			EndIndex: 136,
		},
		helpers.TabField{
			Name: "Claim IDR Load Date",
			StartIndex: 137,
			EndIndex: 146,
		},
		helpers.TabField{
			Name: "Claim Control Number",
			StartIndex: 147,
			EndIndex: 186,
		},
		helpers.TabField{
			Name: "Beneficiary Equitable BIC HICN Number",
			StartIndex: 187,
			EndIndex: 197,
		},
		helpers.TabField{
			Name: "Claim Line Allowed Charges Amount",
			StartIndex: 198,
			EndIndex: 214,
		},
		helpers.TabField{
			Name: "Claim Disposition Code",
			StartIndex: 215,
			EndIndex: 216,
		},
	},
	"usgov.hhs.cclf_partd": []helpers.TabField{
		helpers.TabField{
			Name: "Current Claim Unique Identifier",
			StartIndex: 1,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 14,
			EndIndex: 24,
		},
		helpers.TabField{
			Name: "NDC Code",
			StartIndex: 25,
			EndIndex: 35,
		},
		helpers.TabField{
			Name: "Claim Type Code",
			StartIndex: 36,
			EndIndex: 37,
		},
		helpers.TabField{
			Name: "Claim Line From Date",
			StartIndex: 38,
			EndIndex: 47,
		},
		helpers.TabField{
			Name: "Provider Service Identifier Qualifier Code",
			StartIndex: 48,
			EndIndex: 49,
		},
		helpers.TabField{
			Name: "Claim Service Provider Generic ID Number",
			StartIndex: 50,
			EndIndex: 69,
		},
		helpers.TabField{
			Name: "Claim Dispensing Status Code",
			StartIndex: 70,
			EndIndex: 70,
		},
		helpers.TabField{
			Name: "Claim Dispense as Written (DAW) Product Selection Code",
			StartIndex: 71,
			EndIndex: 71,
		},
		helpers.TabField{
			Name: "Claim Line Service Unit Quantity",
			StartIndex: 72,
			EndIndex: 84,
		},
		helpers.TabField{
			Name: "Claim Line Days Supply Quantity",
			StartIndex: 85,
			EndIndex: 93,
		},
		helpers.TabField{
			Name: "Provider Prescribing ID Qualifier Code",
			StartIndex: 94,
			EndIndex: 95,
		},
		helpers.TabField{
			Name: "Claim Prescribing Provider Generic ID Number",
			StartIndex: 96,
			EndIndex: 115,
		},
		helpers.TabField{
			Name: "Claim Line Beneficiary Payment Amount",
			StartIndex: 116,
			EndIndex: 128,
		},
		helpers.TabField{
			Name: "Claim Adjustment Type Code",
			StartIndex: 129,
			EndIndex: 130,
		},
		helpers.TabField{
			Name: "Claim Effective Date",
			StartIndex: 131,
			EndIndex: 140,
		},
		helpers.TabField{
			Name: "Claim IDR Load Date",
			StartIndex: 141,
			EndIndex: 150,
		},
		helpers.TabField{
			Name: "Claim Line Prescription Service Reference Number",
			StartIndex: 151,
			EndIndex: 162,
		},
		helpers.TabField{
			Name: "Claim Line Prescription Fill Number",
			StartIndex: 163,
			EndIndex: 171,
		},
	},
	"usgov.hhs.cclf_demographics": []helpers.TabField{
		helpers.TabField{
			Name: "Beneficiary HIC Number",
			StartIndex: 1,
			EndIndex: 11,
		},
		helpers.TabField{
			Name: "Beneficiary FIPS State Code",
			StartIndex: 12,
			EndIndex: 13,
		},
		helpers.TabField{
			Name: "Beneficiary FIPS County Code",
			StartIndex: 14,
			EndIndex: 16,
		},
		helpers.TabField{
			Name: "Beneficiary ZIP Code",
			StartIndex: 17,
			EndIndex: 21,
		},
		helpers.TabField{
			Name: "Beneficiary Date of Birth",
			StartIndex: 22,
			EndIndex: 31,
		},
		helpers.TabField{
			Name: "Beneficiary Sex Code",
			StartIndex: 32,
			EndIndex: 32,
		},
		helpers.TabField{
			Name: "Beneficiary Race Code",
			StartIndex: 33,
			EndIndex: 33,
		},
		helpers.TabField{
			Name: "Beneficiary Age",
			StartIndex: 34,
			EndIndex: 36,
		},
		helpers.TabField{
			Name: "Beneficiary Medicare Status Code",
			StartIndex: 37,
			EndIndex: 38,
		},
		helpers.TabField{
			Name: "Beneficiary Dual Status Code",
			StartIndex: 39,
			EndIndex: 40,
		},
		helpers.TabField{
			Name: "Beneficiary Death Date",
			StartIndex: 41,
			EndIndex: 50,
		},
		helpers.TabField{
			Name: "Date beneficiary enrolled in Hospice",
			StartIndex: 51,
			EndIndex: 60,
		},
		helpers.TabField{
			Name: "Date beneficiary ended Hospice",
			StartIndex: 61,
			EndIndex: 70,
		},
		helpers.TabField{
			Name: "Beneficiary First Name",
			StartIndex: 71,
			EndIndex: 100,
		},
		helpers.TabField{
			Name: "Beneficiary Middle Name",
			StartIndex: 101,
			EndIndex: 115,
		},
		helpers.TabField{
			Name: "Beneficiary Last Name",
			StartIndex: 116,
			EndIndex: 155,
		},
	},
	"usgov.hhs.cclf_bene_xref": []helpers.TabField{
		helpers.TabField{
			Name: "Current HIC Number",
			StartIndex: 1,
			EndIndex: 11,
		},
		helpers.TabField{
			Name: "Previous HIC Number",
			StartIndex: 12,
			EndIndex: 22,
		},
		helpers.TabField{
			Name: "Previous HICN Effective Date",
			StartIndex: 23,
			EndIndex: 32,
		},
		helpers.TabField{
			Name: "Previous HICN Obsolete Date",
			StartIndex: 33,
			EndIndex: 42,
		},
		helpers.TabField{
			Name: "Beneficiary Railroad Board Number",
			StartIndex: 43,
			EndIndex: 54,
		},
	},
}

func (d *Description) Available() ([]bloomsource.Source, error) {
	sources := []bloomsource.Source{}

	cclfDir := viper.GetString("cclfDir")

	files, err := ioutil.ReadDir(cclfDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
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
					})
				}
			}
		}
	}

	return sources, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	fields := sourceFields[sourceName]
	names := make([]string, len(fields))

	for i, field := range fields {
		names[i] = field.Name
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

	files, err := ioutil.ReadDir(cclfDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if fileRegex.MatchString(file.Name()) {
				fullPath := viper.GetString("cclfDir") + "/" + file.Name()
				file, err := os.Open(fullPath)
				if err != nil {
					return nil, err
				}

				tabReader := helpers.NewTabReader(file, sourceFields[source.Name])

				return tabReader, nil
			}
		}
	}

	return nil, errors.New("Source file not found " + source.Name + " " + source.Version)
}