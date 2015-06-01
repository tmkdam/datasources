package main

import (
	"database/sql"
	"github.com/gocodo/bloomdb"
	"github.com/spf13/viper"
	"encoding/csv"
	"os"
	"math/rand"
	"strings"
	"reflect"
	"strconv"
	"time"

	"fmt"
)

type tableColumnInfo struct {
	Name string
	Type string
}

func asString(src interface{}) string {
	if src == nil {
		return ""
	}

	switch v := src.(type) {
		case string:
			return v
		case time.Time:
			return v.Format("2006-01-02")
		case []byte:
			return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(rv.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(rv.Uint(), 10)
		case reflect.Float64:
			return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
		case reflect.Float32:
			return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
		case reflect.Bool:
			return strconv.FormatBool(rv.Bool())
	}

	return fmt.Sprintf("%v", src)
}

func tableColumns(conn *sql.DB, table string) ([]tableColumnInfo, error) {
	columns := []tableColumnInfo{}
	rows, err := conn.Query(`	SELECT column_name, data_type 
														FROM information_schema.columns
													 	WHERE table_name = '` + table + `';`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, columnType string
		if err := rows.Scan(&name, &columnType); err != nil {
			return nil, err
		}

		columns = append(columns, tableColumnInfo{
				name,
				columnType,
			})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

func LookupColumn(columns []string, contents []interface{}, columnName string) interface{} {
	var index int
	for i, column := range columns {
		if column == columnName {
			index = i
			break
		}
	}

	return asString(contents[index])
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
func GenerateUid() string {
	b := make([]rune, 15)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	configPath := os.Getenv("BLOOM_CONFIG")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uidFile, err := os.Create("uid_to_hicn.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer uidFile.Close()
	uidWriter := csv.NewWriter(uidFile)

	uidWriter.Write([]string{ "uid", "hicn" })

	beneFile, err := os.Create("uid_beneficiaries.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer beneFile.Close()
	beneWriter := csv.NewWriter(beneFile)

	bloomdb := bloomdb.CreateDB()

	conn, err := bloomdb.SqlConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select all XRefs and build hash of each HICN to list of related HICN
	relatedHicns := map[string][]string{}
	rows, err := conn.Query("SELECT current_hic_number, previous_hic_number FROM usgov_hhs_cclf_bene_xrefs")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var currHicn, prevHicn string
		if err := rows.Scan(&currHicn, &prevHicn); err != nil {
			fmt.Println(err)
			return
		}

		currList, currOk := relatedHicns[currHicn]
		prevList, prevOk := relatedHicns[prevHicn]


		if !currOk {
			currList = []string{}
		}

		if !prevOk {
			prevList = []string{}
		}

		newCurrList := append(currList, prevList...)
		newPrevList := append(prevList, currList...)

		newCurrList = append(currList, prevHicn)
		newPrevList = append(prevList, currHicn)

		relatedHicns[currHicn] = newCurrList
		relatedHicns[prevHicn] = newPrevList
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Get beneficiary details
	columns := []string{}
	uidColumns := []string{}
	infos, err := tableColumns(conn, "usgov_hhs_cclf_beneficiaries")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range infos {
		if info.Type != "uuid" && !strings.HasPrefix(info.Name, "bloom_") && info.Name != date_beneficiary_enrolled_in_hospice {
			columns = append(columns, info.Name)
			if info.Name != "beneficiary_hic_number" {
				uidColumns = append(uidColumns, info.Name)
			} else {
				uidColumns = append(uidColumns, "uid")
			}
		}
	}

	beneWriter.Write(uidColumns)

	// For each Benificiary
	relatedUids := map[string]string{}
	rows, err = conn.Query("SELECT " + strings.Join(columns, ",") + " FROM usgov_hhs_cclf_beneficiaries")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		resultRefs := make([]interface{}, len(columns))
		results := make([]interface{}, len(columns))

		for i, _ := range columns {
			resultRefs[i] = &results[i]
		}

		if err := rows.Scan(resultRefs...); err != nil {
			fmt.Println(err)
			return
		}

		// if no related HICN
		// - create related HICN with empty list
		hcin := asString(LookupColumn(columns, results, "beneficiary_hic_number"))
		hicnList, ok := relatedHicns[hcin]
		if !ok {
			relatedHicns[hcin] = []string{}
		}

		// if UID not already created
		// - output beneficiray + new UID
		// else
		// - Generate UID
		uid, ok := relatedUids[hcin]
		if !ok {
			uid = GenerateUid()

			for _, relatedHicn := range hicnList {
				relatedUids[relatedHicn] = uid
				uidWriter.Write([]string{ uid, relatedHicn })
			}
			
			relatedUids[hcin] = uid
			uidWriter.Write([]string{ uid, hcin })
		}

		uidBene := make([]string, len(columns))
		for i, col := range columns {
			if col != "beneficiary_hic_number" {
				uidBene[i] = asString(results[i])
			} else {
				uidBene[i] = uid
			}
		}

		beneWriter.Write(uidBene)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	uidWriter.Flush()
	beneWriter.Flush()
}