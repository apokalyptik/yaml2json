package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func init() {
}

func main() {
	var data map[string]interface{}
	if raw, err := ioutil.ReadAll(os.Stdin); err != nil {
		log.Fatalf("Error reading stdin: %s", err.Error())
	} else {
		if err := yaml.Unmarshal(raw, &data); err != nil {
			log.Fatalf("Error unmarshaling YAML: %s", err.Error())
		}
	}
	if raw, err := json.MarshalIndent(data, "", "\t"); err != nil {
		log.Fatalf("Error encoding JSON: %s", err.Error())
	} else {
		os.Stdout.Write(raw)
	}
}
