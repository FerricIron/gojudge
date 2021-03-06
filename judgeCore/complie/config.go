package complie

import (
	"encoding/json"
	"io/ioutil"
)

var configPath = "/home/gojudge/judgeCore/config.json"
type configStruct map[string]struct {
	Argv      []string `json:"argv"`
	TimeLimit int      `json:"timelimit"`
}

var Config configStruct

func ParseConfig() error {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &Config)
}
