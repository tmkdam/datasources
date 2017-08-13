package main

import (
	"log"
	"strings"
	"os/exec"
	"bytes"
	"io/ioutil"
	"os"
)

var ignoreList = []string {
	"github.com/bloomapi/datasources/runner",
	"github.com/bloomapi/datasources/usgov/hhs/cclf",
	"github.com/bloomapi/datasources/usgov/hhs/ahrq/md_sid_2011",
	"github.com/bloomapi/datasources/usgov/hhs/hcris",
	"github.com/bloomapi/datasources/usgov/hhs/phyreferral",
	"github.com/bloomapi/datasources/usgov/hhs/hosp_comp",
	"github.com/bloomapi/datasources/usgov/hhs/mu_attestation",
	"github.com/bloomapi/datasources/usgov/hhs/mu_report",
	"github.com/bloomapi/datasources/usgov/hhs/partd_utilization",
	"github.com/bloomapi/datasources/usgov/hhs/mu_eh_attestation",
	"github.com/bloomapi/datasources/usgov/hhs/mu_eh_payments",
}

func searchDir(path string) []string {
	command := exec.Command("go", "list")
	command.Dir = path
	out := bytes.Buffer{}
	command.Stdout = &out

	if err := command.Run(); err != nil {
		files, err := ioutil.ReadDir(path)
		if (err != nil) {
			log.Fatal(err)
		}

		packages := []string{}
		for _, file := range files {
			if file.IsDir() {
				dirPackages := searchDir(path + "/" + file.Name())
				packages = append(packages, dirPackages...)
			}
		}	

		return packages
	}

	output := string(out.Bytes())
	output = strings.Trim(output, "\n")

	for _, ignore := range ignoreList {
		if output == ignore {
			return []string{}
		}
	}

	return []string{ strings.Trim(output, "\n") }
}

func main() {
	files, err := ioutil.ReadDir("./")
	if (err != nil) {
		log.Fatal(err)
	}

	packages := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			dirPackages := searchDir("./" + file.Name())
			packages = append(packages, dirPackages...)
		}
	}

	file, err := os.Create("./Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("export PATH := /usr/local/gonative/go/bin:$(PATH)\n\n")
	file.WriteString("GOPATH ?= $(HOME)/go\n\n")
	file.WriteString("all:\n")
	file.WriteString("\tmkdir -p $(GOPATH)/bin/linux_amd64/datasources/\n")
	file.WriteString("\tcp $(GOPATH)/src/github.com/bloomapi/datasources/config.toml $(GOPATH)/bin/linux_amd64/datasources\n")
	file.WriteString("\tenv GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/linux_amd64/datasources/runner github.com/bloomapi/datasources/runner\n")

	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/bloomapi/datasources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		file.WriteString("\tmkdir -p /$(GOPATH)/bin/linux_amd64/datasources/sources/" + folderName + "\n")
		file.WriteString("\tenv GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/linux_amd64/datasources/sources/" + folderName + "/" + execName + " " + packageName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/dbmapping.yaml $(GOPATH)/bin/linux_amd64/datasources/sources/" + folderName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/searchmapping.yaml $(GOPATH)/bin/linux_amd64/datasources/sources/" + folderName + "\n")
	}

	file.WriteString("\n")
	file.WriteString("native:\n")
	file.WriteString("\tmkdir -p /$(GOPATH)/bin/datasources/\n")
	file.WriteString("\tcp $(GOPATH)/src/github.com/bloomapi/datasources/config.toml $(GOPATH)/bin/datasources\n")
	file.WriteString("\tgo build -o $(GOPATH)/bin/datasources/runner github.com/bloomapi/datasources/runner\n")
	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/bloomapi/datasources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		file.WriteString("\tmkdir -p $(GOPATH)/bin/datasources/sources/" + folderName + "\n")
		file.WriteString("\tgo build -o $(GOPATH)/bin/datasources/sources/" + folderName + "/" + execName + " " + packageName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/dbmapping.yaml $(GOPATH)/bin/datasources/sources/" + folderName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/searchmapping.yaml $(GOPATH)/bin/datasources/sources/" + folderName + "\n")
	}

	file.WriteString(`
clean:
	rm -rf $(GOPATH)/bin/linux_amd64/datasources
	rm -rf $(GOPATH)/bin/datasources

`)

	configFile, err := os.Create("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configFile.WriteString(`sqlConnStr = "postgres://postgres@db/bloomapi?sslmode=disable"
searchHosts = [
  "search"
]
sourceBinaries = [
`)

	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/bloomapi/datasources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		configFile.WriteString("\"./sources/" + folderName + "/" + execName + "\",\n")
	}

	configFile.WriteString("]\n")
}
