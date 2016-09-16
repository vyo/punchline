package controller

import (
	"github.com/vyo/punch/model"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func LoadConfig(configDir string) (model.Config, error) {

	var config model.Config
	content, err := ioutil.ReadFile(configDir + "/punch.json")

	if err != nil {
		return config, err
	} else {
		err = json.Unmarshal(content, config)
		return config, err
	}
}

func StoreConfig(config model.Config) error {
	content, err := json.Marshal(config)

	if err != nil {
		return err
	} else {
		fmt.Printf("writing to %v\n", (*config.GetDir())["config"] + "/punch.json")
		return ioutil.WriteFile((*config.GetDir())["config"] + "/punch.json", content, 0644)
	}
}
