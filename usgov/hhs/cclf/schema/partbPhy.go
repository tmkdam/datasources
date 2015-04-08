package schema

import "github.com/gocodo/bloomsource/helpers"

var PartbPhy = &Cclf{
	Versions: []string{
		"D000101", // Jan 1, 2000
		"D141001", // Oct 1, 2014
	},
	Schemas: [][]helpers.TabField{ // Jan 1, 2000
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
		[]helpers.TabField{ // Oct 1, 2014
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
				Name: "Claim Line Service Unit Quantity",
				StartIndex: 229,
				EndIndex: 252,
			},
			helpers.TabField{
				Name: "HCPCS First Modifier Code",
				StartIndex: 253,
				EndIndex: 254,
			},
			helpers.TabField{
				Name: "HCPCS Second Modifier Code",
				StartIndex: 255,
				EndIndex: 256,
			},
			helpers.TabField{
				Name: "HCPCS Third Modifier Code",
				StartIndex: 257,
				EndIndex: 258,
			},
			helpers.TabField{
				Name: "HCPCS Fourth Modifier Code",
				StartIndex: 259,
				EndIndex: 260,
			},
			helpers.TabField{
				Name: "HCPCS Fifth Modifier Code",
				StartIndex: 261,
				EndIndex: 262,
			},
			helpers.TabField{
				Name: "Claim Disposition Code",
				StartIndex: 263,
				EndIndex: 264,
			},
			helpers.TabField{
				Name: "Claim Diagnosis First Code",
				StartIndex: 265,
				EndIndex: 271,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Second Code",
				StartIndex: 272,
				EndIndex: 278,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Third Code",
				StartIndex: 279,
				EndIndex: 285,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Fourth Code",
				StartIndex: 286,
				EndIndex: 292,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Fifth Code",
				StartIndex: 293,
				EndIndex: 299,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Sixth Code",
				StartIndex: 300,
				EndIndex: 306,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Seventh Code",
				StartIndex: 307,
				EndIndex: 313,
			},
			helpers.TabField{
				Name: "Claim Diagnosis Eighth Code",
				StartIndex: 314,
				EndIndex: 320,
			},
		},
	},
}