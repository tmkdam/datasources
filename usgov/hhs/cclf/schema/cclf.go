package schema

import (
	"sort"
	"bitbucket.org/gocodo/bloomsource/helpers"
)

type Cclf struct {
	Versions []string
	Schemas [][]helpers.TabField
}

func (c *Cclf) FieldNames() []string {
	fieldNames := []string{}
	seen := map[string]bool{}
	for _, schema := range c.Schemas {
		for _, field := range schema {
			if _, ok := seen[field.Name]; !ok {
				fieldNames = append(fieldNames, field.Name)
				seen[field.Name] = true
			}
		}
	}

	return fieldNames
}

func (c *Cclf) Fields(lookupVersion string) []helpers.TabField {
	versionIndex := 0

	versions := c.Versions
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	for i, schemaVersion := range versions {
		if schemaVersion < lookupVersion {
			versionIndex = len(c.Versions) - i - 1
			break;
		}
	}

	baseSchema := c.Schemas[versionIndex]
	schema := make([]helpers.TabField, len(baseSchema))
	copy(schema, baseSchema)

	// append dummy fields
	fieldNames := c.FieldNames()
	missingFields := []string{}
	for _, fieldName := range fieldNames {
		found := false
		for _, schemaField := range schema {
			if schemaField.Name == fieldName {
				found = true
				break
			}
		}

		if !found {
			missingFields = append(missingFields, fieldName)
		}
	}

	for _, missingField := range missingFields {
		schema = append(schema, helpers.TabField{
			Name: missingField,
			StartIndex: 999999,
			EndIndex: 999999,
		})
	}

	return schema
}