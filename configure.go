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
	"github.com/gocodo/bloomsources/runner",
	"github.com/gocodo/bloomsources/usgov/hhs/cclf",
	"github.com/gocodo/bloomsources/usgov/hhs/hcris",
	"github.com/gocodo/bloomsources/usgov/hhs/hosp_comp",
	"github.com/gocodo/bloomsources/usgov/hhs/mu_attestation",
	"github.com/gocodo/bloomsources/usgov/hhs/mu_report",
	"github.com/gocodo/bloomsources/usgov/hhs/pos",
	"github.com/gocodo/bloomsources/usgov/hhs/ahrq/md_sid_2011",
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

	file.WriteString("all:\n")
	file.WriteString("\tmkdir -p $(GOPATH)/bin/linux_amd64/bloomsources/\n")
	file.WriteString("\tcp $(GOPATH)/src/github.com/gocodo/bloomsources/config.toml $(GOPATH)/bin/linux_amd64/bloomsources\n")
	file.WriteString("\tgox -osarch=\"linux/amd64\" -output $(GOPATH)/bin/linux_amd64/bloomsources/runner github.com/gocodo/bloomsources/runner\n")

	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/gocodo/bloomsources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		file.WriteString("\tmkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/sources/" + folderName + "\n")
		file.WriteString("\tgox -osarch=\"linux/amd64\" -output $(GOPATH)/bin/linux_amd64/bloomsources/sources/" + folderName + "/" + execName + " " + packageName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/sources/" + folderName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/sources/" + folderName + "\n")
	}

	file.WriteString("\n")
	file.WriteString("native:\n")
	file.WriteString("\tmkdir -p /$(GOPATH)/bin/bloomsources/\n")
	file.WriteString("\tcp $(GOPATH)/src/github.com/gocodo/bloomsources/config.toml $(GOPATH)/bin/bloomsources\n")
	file.WriteString("\tgo build -o $(GOPATH)/bin/bloomsources/runner github.com/gocodo/bloomsources/runner\n")
	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/gocodo/bloomsources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		file.WriteString("\tmkdir -p $(GOPATH)/bin/bloomsources/sources/" + folderName + "\n")
		file.WriteString("\tgo build -o $(GOPATH)/bin/bloomsources/sources/" + folderName + "/" + execName + " " + packageName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/dbmapping.yaml $(GOPATH)/bin/bloomsources/sources/" + folderName + "\n")
		file.WriteString("\tcp $(GOPATH)/src/" + packageName + "/searchmapping.yaml $(GOPATH)/bin/bloomsources/sources/" + folderName + "\n")
	}

	file.WriteString(`
clean:
	rm -rf $(GOPATH)/bin/linux_amd64/bloomsources
	rm -rf $(GOPATH)/bin/bloomsources

`)

	configFile, err := os.Create("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configFile.WriteString(`sqlConnStr = "postgres://localhost/bloomapi?sslmode=disable"
searchHosts = [
  "localhost"
]
sourceBinaries = [
`)

	for _, packageName := range packages {
		folderName := strings.Replace(packageName, "github.com/gocodo/bloomsources/", "", 1)
		folderName = strings.Replace(folderName, "/", ".", -1)
		parts := strings.Split(packageName, "/")
		execName := parts[len(parts) - 1]
		configFile.WriteString("\"./sources/" + folderName + "/" + execName + "\",\n")
	}

	configFile.WriteString("]\n")
}
