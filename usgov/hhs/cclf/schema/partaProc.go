package schema

import "github.com/bloomapi/dataloading/helpers"

var PartaProc = &Cclf{
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
	},
}