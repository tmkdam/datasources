package schema

import "github.com/bloomapi/dataloading/helpers"

var Beneficiary = &Cclf{
	Versions: []string{
		"D000101", // Jan 1, 2000
		"D141001", // Oct 1, 2014
	},
	Schemas: [][]helpers.TabField{
		[]helpers.TabField{
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
		[]helpers.TabField{ // Oct 1, 2014
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
			// New As of Oct 1, 2014 v8.0
			helpers.TabField{
				Name: "Beneficiary Original Entitlement Reason Code",
				StartIndex: 156,
				EndIndex: 156,
			},
			helpers.TabField{
				Name: "Beneficiary Entitlement Buy-in Indicator",
				StartIndex: 157,
				EndIndex: 157,
			},
		},
	},
}