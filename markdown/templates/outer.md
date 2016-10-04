# {{.Info.Title}} v{{.Info.Version}} 
{{.Info.Description}}

* Host `{{.Host}}`
* Base Path `{{.BasePath}}`

{{range $uri, $path := .Paths}}
{{with $op := $path.get}}
## GET {{$uri}}
{{template "inner.md" .}}
{{end}}
{{end}}
