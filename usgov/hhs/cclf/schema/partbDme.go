package schema

import "bitbucket.org/gocodo/bloomsource/helpers"

var PartbDme = &Cclf{
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
	},
}