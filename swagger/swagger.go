package swagger

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader) (*API, error) {
	var api API
	err := json.NewDecoder(r).Decode(&api)
	return &api, err
}

type API struct {
	Swagger  string             `json:"swagger"`
	Info     Info               `json:"info"`
	Host     string             `json:"host"`
	BasePath string             `json:"basePath"`
	Schemes  []string           `json:"schemes"`
	Consumes []string           `json:"consumes"`
	Produces []string           `json:"produces"`
	Paths    map[string]Methods `json:"paths"`
}

type Methods struct {
	Get    *Operation `json:"get"`
	Post   *Operation `json:"post"`
	Put    *Operation `json:"put"`
	Patch  *Operation `json:"patch"`
	Delete *Operation `json:"delete"`
	Head   *Operation `json:"head"`
}

type Operation struct {
	Description string              `json:"description"`
	OperationID string              `json:"operationId"`
	Parameters  []Parameter         `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
}

type Response struct {
	Description string `json:"description"`
	Schema      Schema `json:"schema"`
}

type Schema struct {
	Type  string                 `json:"type"`
	Items map[string]interface{} `json:"items"`
}

type Parameter struct {
	Name             string                 `json:"name"`
	In               string                 `json:"in"`
	Description      string                 `json:"description"`
	Required         bool                   `json:"required"`
	Type             string                 `json:"type"`
	CollectionFormat string                 `json:"collectionFormat"`
	Items            map[string]interface{} `json:"items"`
}

type Info struct {
	Version        string  `json:"version"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
