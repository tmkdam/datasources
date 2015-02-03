package main

import (
	"os"
	"io"
	"strings"
	"log"
	"encoding/csv"
)

func main() {
	file, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	specialtyFile, err := os.Create("medicare_specialty_codes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer specialtyFile.Close()
	specialtyWriter := csv.NewWriter(specialtyFile)
	err = specialtyWriter.Write([]string{ "medicare_specialty_code", "medicare_specialty_description" })
	if err != nil {
		log.Fatal(err)
	}

	taxonomyFile, err := os.Create("nucc_taxonomy_codes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer taxonomyFile.Close()
	taxonomyWriter := csv.NewWriter(taxonomyFile)
	err = taxonomyWriter.Write([]string{ "taxonomy_code", "taxonomy_description" })
	if err != nil {
		log.Fatal(err)
	}

	crosswalkFile, err := os.Create("medicare_specialty_to_nucc_taxonomy.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer crosswalkFile.Close()
	crosswalkWriter := csv.NewWriter(crosswalkFile)
	err = crosswalkWriter.Write([]string{ "medicare_specialty_code", "nucc_taxonomy_code" })
	if err != nil {
		log.Fatal(err)
	}

	specialtySeen := map[string]bool{}
	taxonomySeen := map[string]bool{}

	reader := csv.NewReader(file)
	// Read headers, expects:
	// 0: medicare_specialty_code
	// 1: medicare_specialty_description
	// 2: taxonomy_code
	// 3: taxonomy_description
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		specialtyCode := strings.Trim(record[0], " ")
		specialtyDesc := strings.Trim(record[1], " ")
		taxonomyCode := strings.Trim(record[2], " ")
		taxonomyDesc := strings.Trim(record[3], " ")

		if specialtyCode != "" && !specialtySeen[specialtyCode] {
			specialtySeen[specialtyCode] = true
			err = specialtyWriter.Write([]string{specialtyCode, specialtyDesc})
			if err != nil {
				log.Fatal(err)
			}
		}

		if taxonomyCode != "" && !taxonomySeen[taxonomyCode] {
			taxonomySeen[taxonomyCode] = true
			err = taxonomyWriter.Write([]string{taxonomyCode, taxonomyDesc})
			if err != nil {
				log.Fatal(err)
			}
		}

		if taxonomyCode != "" && specialtyCode != "" {
			err = crosswalkWriter.Write([]string{specialtyCode, taxonomyCode})
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	specialtyWriter.Flush()
	taxonomyWriter.Flush()
	crosswalkWriter.Flush()
}