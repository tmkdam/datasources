package schema

import "github.com/gocodo/bloomsource/helpers"

var Partd = &Cclf{
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
		[]helpers.TabField{ // Oct 1, 2014
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
				EndIndex: 95,
			},
			helpers.TabField{
				Name: "Claim Line Days Supply Quantity",
				StartIndex: 96,
				EndIndex: 104,
			},
			helpers.TabField{
				Name: "Provider Prescribing ID Qualifier Code",
				StartIndex: 105,
				EndIndex: 106,
			},
			helpers.TabField{
				Name: "Claim Prescribing Provider Generic ID Number",
				StartIndex: 107,
				EndIndex: 126,
			},
			helpers.TabField{
				Name: "Claim Line Beneficiary Payment Amount",
				StartIndex: 127,
				EndIndex: 139,
			},
			helpers.TabField{
				Name: "Claim Adjustment Type Code",
				StartIndex: 140,
				EndIndex: 141,
			},
			helpers.TabField{
				Name: "Claim Effective Date",
				StartIndex: 142,
				EndIndex: 151,
			},
			helpers.TabField{
				Name: "Claim IDR Load Date",
				StartIndex: 152,
				EndIndex: 161,
			},
			helpers.TabField{
				Name: "Claim Line Prescription Service Reference Number",
				StartIndex: 162,
				EndIndex: 173,
			},
			helpers.TabField{
				Name: "Claim Line Prescription Fill Number",
				StartIndex: 174,
				EndIndex: 182,
			},
		},
	},
}