package schema

import "github.com/gocodo/bloomsource/helpers"

var PartaRevenue = &Cclf{
	Versions: []string{
		"D000101", // Jan 1, 2000
		"D141001", // Oct 1, 2014
	},
	Schemas: [][]helpers.TabField{ // Version Jan 1, 2000
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
		[]helpers.TabField{ // Version Oct 1, 2014
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
			// Changes start here
			helpers.TabField{
				Name: "Claim Line Service Unit Quantity",
				StartIndex: 113,
				EndIndex: 136,
			},
			helpers.TabField{
				Name: "Claim Line Covered Paid Amount", // This field replaced "Claim Line Allowed Unit Quantity"
				StartIndex: 137,
				EndIndex: 153,
			},
			helpers.TabField{
				Name: "HCPCS First Modifier Code",
				StartIndex: 154,
				EndIndex: 155,
			},
			helpers.TabField{
				Name: "HCPCS Second Modifier Code",
				StartIndex: 156,
				EndIndex: 157,
			},
			helpers.TabField{
				Name: "HCPCS Third Modifier Code",
				StartIndex: 158,
				EndIndex: 159,
			},
			helpers.TabField{
				Name: "HCPCS Fourth Modifier Code",
				StartIndex: 160,
				EndIndex: 161,
			},
			helpers.TabField{
				Name: "HCPCS Fifth Modifier Code",
				StartIndex: 162,
				EndIndex: 163,
			},
		},
	},
}