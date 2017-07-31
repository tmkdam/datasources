package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"os"
	"strings"
	"errors"
	"github.com/bloomapi/dataloading"
	"github.com/bloomapi/dataloading/helpers"
	"strconv"
	"bufio"

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

func (d *Description) Available() ([]dataloading.Source, error) {
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

	return []dataloading.Source{
		dataloading.Source{
			Name: "usgov.hhs.pecos",
			Version: version,
		},
		dataloading.Source{
			Name: "usgov.hhs.pecos_pendingreview",
			Version: version,
		},
		dataloading.Source{
			Name: "usgov.hhs.pecos_pmd",
			Version: version,
		},
	}, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
	return columns, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
	var reader dataloading.ValueReader

	switch (source.Name) {
	case "usgov.hhs.pecos":
		downloader := dataloading.NewDownloader("data/", nil)
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
		downloader := dataloading.NewDownloader("data/", nil)
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
	  skiplineReader := bufio.NewReader(nonPhyFileReader)

	  skiplineReader.ReadLine()

		fileReader := io.MultiReader(phyFileReader, strings.NewReader("\n"), skiplineReader)

		reader = helpers.NewCsvReader(fileReader)
	case "usgov.hhs.pecos_pmd":
		downloader := dataloading.NewDownloader("data/", nil)
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