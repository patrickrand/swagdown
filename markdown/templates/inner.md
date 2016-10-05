`{{.OperationID}}`

{{.Description}}

{{with $paths := FilterParameters .Parameters "path"}}
#### Path Parameters
{{range $param := $paths}}
* `{{$param.Name}}` {{$param.Description}}
{{end}}{{end}}
{{with $queries := FilterParameters .Parameters "query"}}
#### Query Parameters
{{range $param := $queries}}
* `{{$param.Name}}` (required={{$param.Required}}) {{$param.Description}}
{{end}}{{end}}
{{with $bodies := FilterParameters .Parameters "body"}}
#### Body Parameters
{{range $param := $bodies}}
* `{{$param.Name}}` (required={{$param.Required}}) {{$param.Description}}
{{end}}
{{end}}
#### Responses 
{{range $code, $resp := .Responses}}
* `{{$code}}` {{$resp.Description}} 
{{end}}