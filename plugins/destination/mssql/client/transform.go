package client

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (*Client) TransformBool(v *schema.Bool) any     { return v.Bool }
func (*Client) TransformBytea(v *schema.Bytea) any   { return v.Bytes }
func (*Client) TransformFloat8(v *schema.Float8) any { return v.Float }
func (*Client) TransformInt8(v *schema.Int8) any     { return v.Int }

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]int64, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Int
	}
	return res
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}
	return string(v.Bytes)
}

func (*Client) TransformUUID(v *schema.UUID) any {
	return mssql.UniqueIdentifier(v.Bytes).String()
}
func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = mssql.UniqueIdentifier(e.Bytes).String()
	}
	return res
}

func transformStringer(v fmt.Stringer) any {
	return v.String()
}

func transformStringerArray[E any](values []E) any {
	res := make([]string, len(values))
	for i, e := range values {
		res[i] = any(&e).(fmt.Stringer).String()
	}
	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) any           { return transformStringer(v) }
func (*Client) TransformCIDRArray(v *schema.CIDRArray) any { return transformStringerArray(v.Elements) }
func (*Client) TransformInet(v *schema.Inet) any           { return transformStringer(v) }
func (*Client) TransformInetArray(v *schema.InetArray) any { return transformStringerArray(v.Elements) }
func (*Client) TransformMacaddr(v *schema.Macaddr) any     { return transformStringer(v) }
func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	return transformStringerArray(v.Elements)
}
func (*Client) TransformText(v *schema.Text) any           { return transformStringer(v) }
func (*Client) TransformTextArray(v *schema.TextArray) any { return transformStringerArray(v.Elements) }
