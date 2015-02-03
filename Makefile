all:
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/runner github.com/gocodo/bloomsources/runner
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/nucc.hcpt/data
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/nucc.hcpt/hcpt github.com/gocodo/bloomsources/nucc/hcpt
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/nucc/hcpt/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/nucc.hcpt
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/nucc/hcpt/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/nucc.hcpt
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.hcpcs/data
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.hcpcs/hcpcs github.com/gocodo/bloomsources/usgov/hhs/hcpcs
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/hcpcs/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.hcpcs
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/hcpcs/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.hcpcs
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.npi/data
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.npi/npi github.com/gocodo/bloomsources/usgov/hhs/npi
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/npi/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.npi
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/npi/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.npi
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.pecos/data
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.pecos/pecos github.com/gocodo/bloomsources/usgov/hhs/pecos
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/pecos/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.pecos
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/pecos/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.pecos
	mkdir -p /$(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.taxonomy_codes/data
	gox -osarch="linux/amd64" -output $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.taxonomy_codes/taxonomy_codes github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes/dbmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.taxonomy_codes
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes/searchmapping.yaml $(GOPATH)/bin/linux_amd64/bloomsources/usgov.hhs.taxonomy_codes

native:
	mkdir -p /$(GOPATH)/bin/bloomsources/
	go build -o $(GOPATH)/bin/bloomsources/runner github.com/gocodo/bloomsources/runner
	mkdir -p /$(GOPATH)/bin/bloomsources/nucc.hcpt/data
	go build -o $(GOPATH)/bin/bloomsources/nucc.hcpt/hcpt github.com/gocodo/bloomsources/nucc/hcpt
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/nucc/hcpt/dbmapping.yaml $(GOPATH)/bin/bloomsources/nucc.hcpt
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/nucc/hcpt/searchmapping.yaml $(GOPATH)/bin/bloomsources/nucc.hcpt
	mkdir -p /$(GOPATH)/bin/bloomsources/usgov.hhs.hcpcs/data
	go build -o $(GOPATH)/bin/bloomsources/usgov.hhs.hcpcs/hcpcs github.com/gocodo/bloomsources/usgov/hhs/hcpcs
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/hcpcs/dbmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.hcpcs
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/hcpcs/searchmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.hcpcs
	mkdir -p /$(GOPATH)/bin/bloomsources/usgov.hhs.npi/data
	go build -o $(GOPATH)/bin/bloomsources/usgov.hhs.npi/npi github.com/gocodo/bloomsources/usgov/hhs/npi
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/npi/dbmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.npi
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/npi/searchmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.npi
	mkdir -p /$(GOPATH)/bin/bloomsources/usgov.hhs.pecos/data
	go build -o $(GOPATH)/bin/bloomsources/usgov.hhs.pecos/pecos github.com/gocodo/bloomsources/usgov/hhs/pecos
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/pecos/dbmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.pecos
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/pecos/searchmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.pecos
	mkdir -p /$(GOPATH)/bin/bloomsources/usgov.hhs.taxonomy_codes/data
	go build -o $(GOPATH)/bin/bloomsources/usgov.hhs.taxonomy_codes/taxonomy_codes github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes/dbmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.taxonomy_codes
	cp $(GOPATH)/src/github.com/gocodo/bloomsources/usgov/hhs/taxonomy_codes/searchmapping.yaml $(GOPATH)/bin/bloomsources/usgov.hhs.taxonomy_codes
