package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) any {
	return v.Bool
}

func (c *Client) TransformBytea(v *schema.Bytea) any {
	return v.Bytes
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (c *Client) TransformCIDR(v *schema.CIDR) any {
	return v.String()
}

func (c *Client) TransformFloat8(v *schema.Float8) any {
	return v.Float
}

func (c *Client) TransformInetArray(v *schema.InetArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (c *Client) TransformInet(v *schema.Inet) any {
	return v.String()
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]int64, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Int
	}
	return res
}

func (c *Client) TransformInt8(v *schema.Int8) any {
	return v.Int
}

func (c *Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}
	return string(v.Bytes)
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) any {
	return v.String()
}

func (c *Client) TransformTextArray(v *schema.TextArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (c *Client) TransformText(v *schema.Text) any {
	return v.String()
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformUUID(v *schema.UUID) any {
	return v.String()
}
