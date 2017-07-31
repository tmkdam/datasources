package schema

import "github.com/bloomapi/dataloading/helpers"

var BeneXref = &Cclf{
	Versions: []string{
		"D000101", // Jan 1, 2000
	},
	Schemas: [][]helpers.TabField{
		[]helpers.TabField{
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
	},
}