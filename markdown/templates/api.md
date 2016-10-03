# {{.Info.Title}} v{{.Info.Version}} 
{{.Info.Description}}

* Host: `{{.Host}}`
* Base Path: `{{.BasePath}}`

{{range $uri, $path := .Paths}}

{{with $op := $path.Get}}
## GET {{$uri}}

{{$op.Description}}

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
{{end}}