{{define "api"}}

# {{.API.Info.Title}} v{{.API.Info.Version}} 
{{.API.Info.Description}}

* Host: `{{.API.Host}}`
* Base Path: `{{.API.BasePath}}`

## GET /api/pets 

Description...

#### Query Parameters

* `tags`: Description...
* `limit`: Description...

#### Responses 
* `200 OK` 

 ```json
 {

 }
 ```


{{end}}