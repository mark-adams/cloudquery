// Code generated by codegen; DO NOT EDIT.

package {{.SubService}}

import (
  "context"

  "github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
  "github.com/cloudquery/plugin-sdk/schema"
  "github.com/PagerDuty/go-pagerduty"
)

func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
  cqClient := meta.(*client.Client)

  more := true
  var offset uint
  for more {
    response, err := cqClient.PagerdutyClient.{{if ne .ListFunctionNameOverride ""}}{{.ListFunctionNameOverride}}{{else}}List{{.StructName}}sWithContext{{end}}(ctx, pagerduty.{{if ne .ListOptionsStructNameOverride ""}}{{.ListOptionsStructNameOverride}}{{else}}List{{.StructName}}sOptions{{end}}{
      Limit: client.MaxPaginationLimit,
      Offset: offset,
    })
    if err != nil {
      return err
    }

    if len(response.{{if ne .ResponseFieldOverride ""}}{{.ResponseFieldOverride}}{{else}}{{.DefaultResponseStructFieldName}}{{end}}) == 0 {
      return nil
    }

    res <- response.{{if ne .ResponseFieldOverride ""}}{{.ResponseFieldOverride}}{{else}}{{.DefaultResponseStructFieldName}}{{end}}

    offset += uint(len(response.{{if ne .ResponseFieldOverride ""}}{{.ResponseFieldOverride}}{{else}}{{.DefaultResponseStructFieldName}}{{end}}))
    more = response.More
  }

  return nil
}