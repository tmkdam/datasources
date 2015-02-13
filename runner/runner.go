package main

import (
	"log"
	"fmt"
	"bytes"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
	"os"
)

func showUsage() {
	fmt.Printf("Usage: %s <command>\n", os.Args[0])
	fmt.Println("=============================\n")
	fmt.Println("Avaialable commands:")
	fmt.Printf("%s bootstrap   # setup datasource in BloomAPI\n", os.Args[0])
	fmt.Printf("%s fetch       # fetch data\n", os.Args[0])
	fmt.Printf("%s update      # fetch/index latest data and add to BloomAPI\n", os.Args[0])
}

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Invalid command usage\n")
		showUsage()
		os.Exit(1)
	}

	arg := os.Args[1]

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	sourceBinaries := viper.GetStringSlice("sourceBinaries")

	switch arg {
	case "update":
		for _, binary := range sourceBinaries {
			parts := strings.Split(binary, "/")
			wd := strings.Join(parts[0:(len(parts) - 1)], "/")
			localBin := "./" + parts[len(parts) - 1]
			command := exec.Command(localBin, "fetch")
			command.Dir = wd
			command.Env = []string{ "BLOOM_CONFIG=" + dir }
			out := bytes.Buffer{}
			command.Stdout = &out
			err = command.Run()
			log.Println(out.String())
			if err != nil {
				log.Println("Error Running Fetch for", binary, err)
			} else {
				log.Println("Successfully Fetched", binary)
			}
		}

		for _, binary := range sourceBinaries {
			parts := strings.Split(binary, "/")
			wd := strings.Join(parts[0:(len(parts) - 1)], "/")
			localBin := "./" + parts[len(parts) - 1]
			command := exec.Command(localBin, "search-index")
			command.Dir = wd
			command.Env = []string{ "BLOOM_CONFIG=" + dir }
			out := bytes.Buffer{}
			command.Stdout = &out
			err = command.Run()
			log.Println(out.String())
			if err != nil {
				log.Println("Error Running Index for", binary, err)
			} else {
				log.Println("Successfully Indexed", binary)
			}
		}

		log.Println("Completed.")
	case "bootstrap":
		for _, binary := range sourceBinaries {
			parts := strings.Split(binary, "/")
			wd := strings.Join(parts[0:(len(parts) - 1)], "/")
			localBin := "./" + parts[len(parts) - 1]
			command := exec.Command(localBin, "bootstrap")
			command.Dir = wd
			command.Env = []string{ "BLOOM_CONFIG=" + dir }
			out := bytes.Buffer{}
			command.Stdout = &out
			err = command.Run()
			log.Println(out.String())
			if err != nil {
				log.Println("Error Running bootstrap for", binary, err)
			}
		}
		log.Println("Completed.")
	case "drop":
		for _, binary := range sourceBinaries {
			parts := strings.Split(binary, "/")
			wd := strings.Join(parts[0:(len(parts) - 1)], "/")
			localBin := "./" + parts[len(parts) - 1]
			command := exec.Command(localBin, "drop")
			command.Dir = wd
			command.Env = []string{ "BLOOM_CONFIG=" + dir }
			out := bytes.Buffer{}
			command.Stdout = &out
			err = command.Run()
			log.Println(out.String())
			if err != nil {
				log.Println("Error Running bootstrap for", binary, err)
			}
		}

		log.Println("Completed.")
	default:
		fmt.Println("Invalid command:", arg)
		showUsage()
		os.Exit(1)
	}
}