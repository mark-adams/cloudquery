// Code generated by codegen; DO NOT EDIT.
package client

import (
    "github.com/cloudquery/cloudquery/plugins/source/datadog/client/services"
)

type DatadogServices struct {
	{{- range $service := . }}
		{{$service.Name}} services.{{$service.ClientName}}
    {{- end }}
}