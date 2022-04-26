package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io"
	"os"
	"strings"
)

func CreateYaml(data interface{}, path string) {
	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("json marshal error: %v", err)
	}
	toYAML, err := yaml.JSONToYAML(marshal)
	if err != nil {
		_ = fmt.Errorf("json to yaml error: %v", err)
	}
	readerYaml := strings.NewReader(string(toYAML))
	f, err := os.Create(path)
	_, err = io.Copy(f, readerYaml)
	if err != nil {
		_ = fmt.Errorf("copy error: %v", err)
	}
}
