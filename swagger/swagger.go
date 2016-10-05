package swagger

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v2"
)

func DecodeJSON(data []byte) (*API, error) {
	var api API
	err := json.Unmarshal(data, &api)
	return &api, err
}

func DecodeYAML(data []byte) (*API, error) {
	var api API
	err := yaml.Unmarshal(data, &api)
	return &api, err
}

func FilterParameters(params []Parameter, t string) []Parameter {
	var results []Parameter
	for _, param := range params {
		if strings.ToLower(param.In) == t {
			results = append(results, param)
		}
	}
	return results
}

type API struct {
	Swagger  string             `json,yaml:"swagger"`
	Info     Info               `json,yaml:"info"`
	Host     string             `json,yaml:"host"`
	BasePath string             `json,yaml:"basePath"`
	Schemes  []string           `json,yaml:"schemes"`
	Consumes []string           `json,yaml:"consumes"`
	Produces []string           `json,yaml:"produces"`
	Paths    map[string]Methods `json,yaml:"paths"`
}

type Methods struct {
	Get    *Operation `json,yaml:"get"`
	Post   *Operation `json,yaml:"post"`
	Put    *Operation `json,yaml:"put"`
	Patch  *Operation `json,yaml:"patch"`
	Delete *Operation `json,yaml:"delete"`
	Head   *Operation `json,yaml:"head"`
}

type Operation struct {
	Description string              `json,yaml:"description"`
	OperationID string              `json,yaml:"operationId"`
	Parameters  []Parameter         `json,yaml:"parameters"`
	Responses   map[string]Response `json,yaml:"responses"`
}

type Response struct {
	Description string `json,yaml:"description"`
	Schema      Schema `json,yaml:"schema"`
}

type Schema struct {
	Type  string                 `json,yaml:"type"`
	Items map[string]interface{} `json,yaml:"items"`
}

type Parameter struct {
	Name             string                 `json,yaml:"name"`
	In               string                 `json,yaml:"in"`
	Description      string                 `json,yaml:"description"`
	Required         bool                   `json,yaml:"required"`
	Type             string                 `json,yaml:"type"`
	CollectionFormat string                 `json,yaml:"collectionFormat"`
	Items            map[string]interface{} `json,yaml:"items"`
}

type Info struct {
	Version        string  `json,yaml:"version"`
	Title          string  `json,yaml:"title"`
	Description    string  `json,yaml:"description"`
	TermsOfService string  `json,yaml:"termsOfService"`
	Contact        Contact `json,yaml:"contact"`
	License        License `json,yaml:"license"`
}

type Contact struct {
	Name  string `json,yaml:"name"`
	Email string `json,yaml:"email"`
	URL   string `json,yaml:"url"`
}

type License struct {
	Name string `json,yaml:"name"`
	URL  string `json,yaml:"url"`
}
