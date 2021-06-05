package loader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func Load(configFile string) (patterns []Pattern, err error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &patterns); err != nil {
		return
	}

	log.Println("Message pattern loaded")

	return
}
