package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"os"
	"strings"
	"errors"
	"bitbucket.org/gocodo/bloomsource"
	"bitbucket.org/gocodo/bloomsource/helpers"
	"strconv"

	"fmt"
)

type Description struct {}

var uris = map[string]string {
	"PECOS": "https://data.cms.gov/api/views/qcn7-gc3g/rows.csv?accessType=DOWNLOAD",
	"PECOS-PhysiciansPendingReview": "https://data.cms.gov/api/views/ip7y-ztn9/rows.csv?accessType=DOWNLOAD",
	"PECOS-NonPhysiciansPendingReview": "https://data.cms.gov/api/views/n86y-dqck/rows.csv?accessType=DOWNLOAD",
	"PECOS-PMD": "https://data.cms.gov/api/views/g6jg-y93m/rows.csv?accessType=DOWNLOAD",
}

var columns = []string{ "npi", "last_name", "first_name" }

func (d *Description) Available() ([]bloomsource.Source, error) {
	resp, err := http.Get("https://data.cms.gov/api/views/qcn7-gc3g")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	info := map[string]interface{}{}
	json.Unmarshal(body, &info)
	version := strconv.FormatFloat(info["viewLastModified"].(float64), 'f', 0, 64)
	fmt.Println(version)

	return []bloomsource.Source{
		bloomsource.Source{
			Name: "usgov.hhs.pecos",
			Version: version,
		},
		bloomsource.Source{
			Name: "usgov.hhs.pecos_pendingreview",
			Version: version,
		},
		bloomsource.Source{
			Name: "usgov.hhs.pecos_pmd",
			Version: version,
		},
	}, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	return columns, nil
}

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
	var reader bloomsource.ValueReader

	switch (source.Name) {
	case "usgov.hhs.pecos":
		downloader := bloomsource.NewDownloader("data/", nil)
	  path, err := downloader.Fetch(uris["PECOS"])
	  if err != nil {
	    return nil, err
	  }

	  fileReader, err := os.Open(path)
	  if err != nil {
	    return nil, err
	  }

	  reader = helpers.NewCsvReader(fileReader)
	case "usgov.hhs.pecos_pendingreview":
		downloader := bloomsource.NewDownloader("data/", nil)
	  phyPath, err := downloader.Fetch(uris["PECOS-PhysiciansPendingReview"])
	  if err != nil {
	    return nil, err
	  }

	  phyFileReader, err := os.Open(phyPath)
	  if err != nil {
	    return nil, err
	  }

	  nonPhyPath, err := downloader.Fetch(uris["PECOS-NonPhysiciansPendingReview"])
	  if err != nil {
	    return nil, err
	  }

	  nonPhyFileReader, err := os.Open(nonPhyPath)
	  if err != nil {
	    return nil, err
	  }

		fileReader := io.MultiReader(phyFileReader, strings.NewReader("\n"), nonPhyFileReader)

		reader = helpers.NewCsvReader(fileReader)
	case "usgov.hhs.pecos_pmd":
		downloader := bloomsource.NewDownloader("data/", nil)
	  path, err := downloader.Fetch(uris["PECOS-PMD"])
	  if err != nil {
	    return nil, err
	  }

	  fileReader, err := os.Open(path)
	  if err != nil {
	    return nil, err
	  }

	  reader = helpers.NewCsvReader(fileReader)
	default:
		return nil, errors.New("Unknown source")
	}

	return reader, nil
}