package schema

import "github.com/gocodo/bloomsource/helpers"

var Parta = &Cclf{
	Versions: []string{
		"D000101", // Jan 1, 2000
	},
	Schemas: [][]helpers.TabField{
		[]helpers.TabField{
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
	},
}