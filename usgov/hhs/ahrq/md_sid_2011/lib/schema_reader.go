package lib

import (
	"bitbucket.org/gocodo/bloomsource/helpers"
	"io"
	"strconv"
	"strings"
	"errors"
)

func buildSchema(name string) ([]helpers.TabField, error) {
	var schema string
	switch (name) {
	case "usgov.hhs.ahrq.md_sid_2011.chgs":
		schema = chgsSchema
	case "usgov.hhs.ahrq.md_sid_2011.core":
		schema = coreSchema
	case "usgov.hhs.ahrq.md_sid_2011.dx_pr_grps":
		schema = dxPrGrpsSchema
	case "usgov.hhs.ahrq.md_sid_2011.severity":
		schema = severitySchema
	}

	schema = strings.Join(strings.Split(schema, "\n")[20:], "\n")

	sReader := strings.NewReader(schema)

  schemaReader := helpers.NewTabReader(sReader, []helpers.TabField{
  		helpers.TabField{
 				Name : "name",
				StartIndex : 30,
				EndIndex : 55,
  		},
  		helpers.TabField{
  			Name : "start",
  			StartIndex : 57,
  			EndIndex : 61,
  		},
  		helpers.TabField{
  			Name : "end",
  			StartIndex : 62,
  			EndIndex : 66,
  		},
  	})

  dataSchema := make([]helpers.TabField, 0)

  for {
  	row, err := schemaReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		name, ok := row.Value("name")
		if !ok {
			return nil, errors.New("Could not find Schema 'name'")
		}

		startIndexRaw, ok := row.Value("start")
		if !ok {
			return nil, errors.New("Could not find Schema 'start'")
		}

		endIndexRaw, ok := row.Value("end")
		if !ok {
			return nil, errors.New("Could not find Schema 'end'")
		}

		startIndex, err := strconv.ParseUint(startIndexRaw, 0, 64)
		if err != nil {
			return nil, err
		}

		endIndex, err := strconv.ParseUint(endIndexRaw, 0 , 64)
		if err != nil {
			return nil, err
		}

		dataSchema = append(dataSchema, helpers.TabField{
				Name: name,
				StartIndex: startIndex,
				EndIndex: endIndex,
			})
  }

  return dataSchema, nil
}