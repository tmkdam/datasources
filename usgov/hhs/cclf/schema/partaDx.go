package schema

import "github.com/gocodo/bloomsource/helpers"

var PartaDx = &Cclf{
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
	},
}