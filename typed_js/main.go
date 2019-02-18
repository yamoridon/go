package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("skimdb.json")
	if err != nil {
		log.Fatal(err)
	}

	var db struct {
		Rows []struct {
			Value struct {
				Name            string                  `json:"name"`
				DevDependencies *map[string]interface{} `json:"devDependencies"`
			} `json:"value"`
		} `json:"rows"`
	}

	if err := json.Unmarshal(bytes, &db); err != nil {
		log.Fatal(err)
	}

	for _, row := range db.Rows {
		var value = row.Value
		var ts, flow bool
		if value.DevDependencies != nil {
			if _, ok := (*value.DevDependencies)["typescript"]; ok {
				ts = true
			}
			if _, ok := (*value.DevDependencies)["flow-bin"]; ok {
				flow = true
			}
		}
		if ts && flow {
			fmt.Printf("%s: TypeScript, Flow\n", value.Name)
		} else if ts {
			fmt.Printf("%s: TypeScript\n", value.Name)
		} else if flow {
			fmt.Printf("%s: Flow\n", value.Name)
		}
	}
}
