# {{.Info.Title}} v{{.Info.Version}} 

{{.Info.Description}}

* Host `{{.Host}}`
* Base Path `{{.BasePath}}`

{{range $uri, $path := .Paths}}
{{with $op := $path.Get}}
## GET {{$uri}}

{{template "inner.md" .}}{{end}}
{{with $op := $path.Post}}
## POST {{$uri}}

{{template "inner.md" .}}{{end}}
{{with $op := $path.Put}}
## PUT {{$uri}}

{{template "inner.md" .}}{{end}}
{{with $op := $path.Patch}}
## PATCH {{$uri}}

{{template "inner.md" .}}{{end}}
{{with $op := $path.Delete}}
## DELETE {{$uri}}

{{template "inner.md" .}}{{end}}
{{with $op := $path.Head}}
## HEAD {{$uri}}

{{template "inner.md" .}}{{end}}{{end}}